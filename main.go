package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New(
		"Bot " + "MTAxNzgxOTY2NDc4MjQ3OTQ0MA.GeFqjV.WaTDeJCOu4xr0JzF8jDIbv1QXG4e3sypvtCwq4",
	)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if strings.Contains(m.Content, "ping") {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if strings.Contains(m.Content, "time") {
		_, err := s.ChannelMessageSend(m.ChannelID, "The current time is: "+time.Now().String())
		if err != nil {
			fmt.Println("error sending time message,", err)
		}
	}
}
