package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Auth token for the bot")
	flag.Parse()
}
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	//good practice is to have the bot ignore its own messages
	if m.Author.ID == s.State.User.ID {
		// si el ID del mensajero es igual al del bot, implicamos que el bot ve un mensaje de el
		return
	}
	if m.Content == "Ping!" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	if m.Content == "Pong!" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}
func main() {
	//Create Discord Bot Session Using the Token
	bot, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Fatal("Error Creating Discord Session,", err)
	}

	//Now we can Add method handlers for it
	bot.AddHandler(MessageCreate)
	//intents are to specify what the Discord bot is gonna do
	//in this case for now, the bot is only going to send messages
	bot.Identify.Intents = discordgo.IntentsGuildMessages

	//Now we do the websocket thing
	err = bot.Open()
	if err != nil {
		log.Fatal("Error Opening Connection, ", err)
	}

	//Now the Program is blocked waiting for it to end
	fmt.Println("Bot is Up!, Press CTRL+C to Exit!.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	//Close Discord Session
	bot.Close()
}
