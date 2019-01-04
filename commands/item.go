package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Pagination struct {
	Page           string
	PageNext       *string
	PagePrev       *string
	PageTotal      string
	Results        string
	ResultsPerPage string
	ResultsTotal   string
}

type Results struct {
	ID      string
	Icon    string
	Name    string
	Url     string
	UrlType string
	_       string
	_Score  string
}

type Response struct {
	Pagination
	Results
}

// GetItemDetails sends a message containing item details from XIVAPI
func GetItemDetails(discord *discordgo.Session, message *discordgo.MessageCreate) {
	url := "https://xivapi.com/search"
	item := "ifrit"
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("string", item)
	q.Add("pretty", "1")
	req.URL.RawQuery = q.Encode()

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	response := Response{}
	// log.Println(response)
	log.Println(req.URL.String())

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(response.Results)
	// discord.ChannelMessageSend(message.ChannelID, "[TEST] did you just type "+message.Content+" ?")
}
