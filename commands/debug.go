package commands

import "github.com/bwmarrin/discordgo"

// DebugMessage is a testing function copying your previous message
func DebugMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	discord.ChannelMessageSend(message.ChannelID, "[TEST] did you just type "+message.Content+" ?")
}
