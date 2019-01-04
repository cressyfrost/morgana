package main

import (
	"fmt"
	"strings"

	"github.com/cressyfrost/morgana/commands"

	"github.com/bwmarrin/discordgo"
)

var (
	commandPrefix string
	botID         string
)

func main() {
	discord, err := discordgo.New("Bot NTI5OTI3MjA5MzQzOTA5ODg4.Dw3_2w.kj7ntGLUw2EJCjDlnQJh4Pt6ATE")
	errCheck("error creating discord session", err)
	user, err := discord.User("@me")
	errCheck("error retrieving account", err)

	botID = user.ID
	discord.AddHandler(commandHandler)
	discord.AddHandler(func(discord *discordgo.Session, ready *discordgo.Ready) {
		err = discord.UpdateStatus(0, "A random not-so-helpful bot.")
		if err != nil {
			fmt.Println("Error attempting to set my status")
		}
		servers := discord.State.Guilds
		fmt.Printf("Morgana has started on %d servers", len(servers))
	})

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	commandPrefix = "!"

	<-make(chan struct{})

}

func errCheck(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %+v", msg, err)
		panic(err)
	}
}

func commandHandler(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	//Do nothing because the bot is talking or message doesn't contain commands
	if user.ID == botID || user.Bot || !strings.HasPrefix(message.Content, commandPrefix) {
		return
	}

	// Get command for routing
	messageContent := strings.Split(strings.ToLower(message.Content), " ")
	command := strings.TrimPrefix(messageContent[0], commandPrefix)

	switch command {
	case "test":
		commands.DebugMessage(discord, message)
	case "item":
		commands.GetItemDetails(discord, message)
	case "weather":
		discord.ChannelMessageSend(message.ChannelID, "Weather is now Sunny")
	default:
		discord.ChannelMessageSend(message.ChannelID, "Hey, i didn't recognize that command!")
	}

	//fmt.Printf("Message: %+v || From: %s\n", message.Message, message.Author)
}
