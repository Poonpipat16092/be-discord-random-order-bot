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
var log = logx.Console()

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}
func main() {
	//Create a new Discord session using the provided bot token.
	session, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Error("error creating Discord session: %s", err.Error())
		return
	}
	// Register the messageCreate func as a callback for MessageCreate events.
	session.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	session.Identify.Intents = discordgo.IntentsGuildMessages

	err = session.Open()
	if err != nil {
		log.Error("error opening connection: %s", err.Error())
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

func messageCreate(session *discordgo.Session, m *discordgo.MessageCreate) {
	log.Info("ID: %s, Content: %s, Channel: %s", m.Author.Username, m.Content, m.ChannelID)
	if m.Author.ID == session.State.User.ID {
		return
	}

	if m.Content == "ping" {
		session.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		session.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
