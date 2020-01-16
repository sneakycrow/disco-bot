package handlers

import "github.com/bwmarrin/discordgo"

func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore all messages by the bot
	if message.Author.ID == session.State.User.ID {
		return
	}
	// if user types ping say pong
	if message.Content == "ping" {
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	}
}