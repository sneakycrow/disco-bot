package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

// MessageCreate listens for a message being created
func MessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	// Ignore all messages by the bot
	if message.Author.ID == session.State.User.ID {
		return
	}
	// if user types 'ping' say pong
	if message.Content == "ping" {
		session.ChannelMessageSend(message.ChannelID, "Pong!")
	}
	// if user types '!kanye' give a kanye quote
	if message.Content == "!kanye" {
		kanyeQuote := getKanyeQuote()
		session.ChannelMessageSend(message.ChannelID, kanyeQuote)
	}
}

// KanyeQuote is a struct for the quiotes from the kanye rest api
type KanyeQuote struct {
	Quote string `json:"quote"`
}

func getKanyeQuote() string {
	// first we query GET the url
	resp, err := http.Get("https://api.kanye.rest")
	// check for error
	if err != nil {
		// return a request error
		return "I couldn't hear kanye"
	}
	// establish our 'k' var which is our KanyeQuote struct
	var k KanyeQuote
	// attempt to decode the response body into our k var
	// note: the k var is actually referencing the pointer address, that way it will modify the var instead of creating a copy
	err = json.NewDecoder(resp.Body).Decode(&k)
	// check again if our err is nil
	if err != nil {
		// return a parse error
		return "I couldn't understand what kanye was saying!"
	}
	// return our kanye quote
	return k.Quote
}
