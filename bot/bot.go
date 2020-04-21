package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	session *discordgo.Session
}

func Start(token string) (*Bot, error) {
	session, err := discordgo.New(fmt.Sprintf("Bot %s", token))
	if err != nil {
		return nil, fmt.Errorf("error initializng session: %w", err)
	}

	err = session.Open()
	if err != nil {
		return nil, fmt.Errorf("error opening connection: %w", err)
	}

	log.Println("Started bot")

	return &Bot{session}, nil
}

func (bot *Bot) Stop() {
	err := bot.session.Close()
	if err != nil {
		log.Printf("error closing session: %s\n", err)
	}

	log.Println("Stopped bot")
}
