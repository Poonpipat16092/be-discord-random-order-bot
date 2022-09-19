package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/PurpleSec/logx"
	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)
var con = logx.Console(logx.Info)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
	con.Info("TOKEN:", Token)
}
func main() {
	//Create a new Discord session using the provided bot token.
	session, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	// Register the messageCreate func as a callback for MessageCreate events.
	session.AddHandler(messageCreate)
	con.Info(session.Token)

	// In this example, we only care about receiving message events.
	session.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = session.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	session.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	con.Info("ID: %s, Content: %s", m.Author.Username, m.Content)
	if m.Author.ID == s.State.User.ID {
		return
	}
	// If the messasge is "ping" reply with "Pong!"
	if m.Content == "ping" {
		con.Info("Ping in")
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	// If the message is "pong" reply with "Ping!"
	if m.Content == "pong" {
		con.Info("pong in")
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
