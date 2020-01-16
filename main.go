package main

import (
	"log" // for logging to console
	"os" // for grabbing the env vars
	"os/signal" // for checking for os signals
	"syscall" // for checking for system calls
	"time" // for spinner animation time

	handlers "github.com/sneakycrow/disco-bot/handlers" // local handler functions

	"github.com/joho/godotenv" // loading env vars from .env file
	"github.com/bwmarrin/discordgo" // golang discord api
	"github.com/briandowns/spinner" // spinner
)

func getDiscordTokenFromEnv() string {
	// Do a token lookup
	// envToken would be the string if it exists in the environment
	// isTokenExist is a boolean that tells us whether the token is undefined in the environment
	envToken, isTokenExist := os.LookupEnv("DISCORD_TOKEN")
	if !isTokenExist {
		// if the token is present (even if it's empty), isToken exist will be true
		log.Fatal("DISCORD_TOKEN is not set\n")
	} else if len(envToken) < 1 {
		// if the token does exist but is empty, we want to fatally log that as well
		log.Fatal("DISCORD_TOKEN is empty\n")
	}
	// return our discord_token
	return envToken
}

func displayStatus(isDisplayed bool) {
	// start or stop depending on bool supplied
	if isDisplayed {
		s.Start()
	} else {
		s.Stop()
	}
}

// create a global spinner variable so we can start and stop from fn init and fn main (also prefix changing from either)
var (
	s = spinner.New(spinner.CharSets[9], 100*time.Millisecond)
)

// init function for initalizing spinner and env vars
func init() {
	// set our prefix to loading for initialization
	s.Prefix = "Loading "
	displayStatus(true)
	// first we load our .env file into the environment
	err := godotenv.Load()

	// check if there was an error loading it
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	// grab our discord token from environment
	discordToken := getDiscordTokenFromEnv()
	// initializers our discord bot
	discordBot, err := discordgo.New("Bot " + discordToken)
	// make sure the bot initialized
	if err != nil {
		log.Fatal("Error initializing Discord Bot")
	}

	// add handler to discord bot
	discordBot.AddHandler(handlers.MessageCreate)
	discordBot.AddHandler(handlers.BotReady)
	// open a websocket connection to discord and begin listening
	err = discordBot.Open()
	// check for error
	if err != nil {
		log.Fatal("error opening connection\nerror: %s\n", err)
	}

	// remove loading display
	displayStatus(false)
	// tell console our bot is running
	log.Printf("Bot is now running. Press CTRL-C to exit.")
	// establish sc channel for detecting os.Signal 1 code
	sc := make(chan os.Signal, 1)
	// have sc be notified on several syscalls and os 
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	// have sc fire off when it receives a signal, which will continue code execution
	<-sc
	// tell spinner prefix to be stopping
	s.Prefix = "Stopping "
	// display stopping status
	displayStatus(true)

	// cleany close down discord session
	discordBot.Close()
	// remove stopping status
	displayStatus(false)
}
