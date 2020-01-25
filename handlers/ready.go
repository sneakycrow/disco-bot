package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

// BotReady is a function that logs the bot ready signal
func BotReady(session *discordgo.Session, websocketReady *discordgo.Ready) {
	if !isChannelIDExist {
		// if it isn't set, log it
		log.Printf("BOT_CHANNEL_ID is not set\n")
	} else if len(botChannelID) < 1 {
		// if it is empty, log it
		log.Printf("BOT_CHANNEL_ID is empty\n")
	} else {
		// let the discord know our bot is ready
		log.Printf("Bot has started successfully")
	}
}
