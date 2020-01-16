package handlers

import (
	"os"
	"fmt"
	
	"github.com/bwmarrin/discordgo"
)

func BotReady(session *discordgo.Session, websocketReady *discordgo.Ready) {
	// Lookup our bot channel id from environment
	botChannelID, isChannelIDExist := os.LookupEnv("BOT_CHANNEL_ID")
	if !isChannelIDExist {
		// if it isn't set, log it
		log.Printf("BOT_CHANNEL_ID is not set\n")
	} else if len(botChannelID) < 1 {
		// if it is empty, log it
		log.Printf("BOT_CHANNEL_ID is empty\n")
	} else {
		// let the discord know our bot is ready
		session.ChannelMessageSend(botChannel, fmt.Sprintf("%s is ready to go", websocketReady.User))
	}
}