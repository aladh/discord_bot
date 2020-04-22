package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/ali-l/discord_bot_go/message"
)

const commandPrefix = "!"

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

func (bot *Bot) AddCommand(command string, handler func(message *message.Message)) {
	bot.session.AddHandler(func(_ *discordgo.Session, msg *discordgo.MessageCreate) {
		if !strings.HasPrefix(commandPrefix+command, msg.Content) {
			return
		}

		handler(message.New(msg.Message, bot))
	})
}

func (bot *Bot) SendMessage(channel string, content string) (*message.Message, error) {
	msg, err := bot.session.ChannelMessageSend(channel, content)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	return message.New(msg, bot), nil
}

func (bot *Bot) EditMessage(channel string, messageID string, content string) error {
	_, err := bot.session.ChannelMessageEdit(channel, messageID, content)
	if err != nil {
		return fmt.Errorf("failed to edit message: %w", err)
	}

	return nil
}
