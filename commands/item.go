package commands

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/spf13/viper"
)

type Pagination struct {
	Page           int
	PageNext       int
	PagePrev       int
	PageTotal      int
	Results        int
	ResultsPerPage int
	ResultsTotal   int
}

type Results struct {
	ID      int
	Icon    string
	Name    string
	Url     string
	UrlType string
	_       string
	_Score  int
}

type Response struct {
	Pagination Pagination
	Results    []Results
	SpeedMs    int
}

// GetItemDetails sends a message containing item details from XIVAPI
func GetItemDetails(discord *discordgo.Session, message *discordgo.MessageCreate) {
	url := "https://xivapi.com/search"
	item := strings.TrimPrefix(message.Content, "!item ")
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
	q.Add("key", viper.GetString("XIVAPI.APIKey"))
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

	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	//var MessageResponse bytes.Buffer
	if (len(response.Results)) > 0 {
		discord.ChannelMessageSend(message.ChannelID, "Displaying top 5 search results for: "+item)
		for key, value := range response.Results {
			// MessageResponse.WriteString(strconv.Itoa(key))
			// MessageResponse.WriteString(". ")
			// MessageResponse.WriteString(value.Name)
			// MessageResponse.WriteString("\n")
			discord.ChannelMessageSend(message.ChannelID, viper.GetString("XIVAPI.BaseURL")+value.Icon)
			discord.ChannelMessageSend(message.ChannelID, "["+strconv.Itoa(value.ID)+"] "+value.Name)
			time.Sleep(200 * time.Millisecond)
			if key >= 4 {
				break
			}
		}
	} else {
		strResponse := "Hmmmmm, I can't find the item you're looking for."
		discord.ChannelMessageSend(message.ChannelID, strResponse)
	}
	// for i := 0; i < v.NumField(); i++ {
	// 	MessageResponse.WriteString(v.Field(i).Interface().(string))
	// 	MessageResponse.WriteString("\n")
	// }

	// if strResponse == "" {
	// 	strResponse = "Hmmmmm, I can't find the item you're looking for."
	// }
	// discord.ChannelMessageSend(message.ChannelID, strResponse)
}
