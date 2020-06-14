package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/ali-l/discord_bot/message"
)

const commandPrefix = "!"

type Bot struct {
	session *discordgo.Session
}

func New(token string) (*Bot, error) {
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("error constructing session: %w", err)
	}

	return &Bot{session}, nil
}

func (bot *Bot) Start() (err error) {
	err = bot.session.Open()
	if err != nil {
		return fmt.Errorf("error opening session: %w", err)
	}
	defer func() { err = bot.Stop() }()

	log.Println("Started bot")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-stop

	return nil
}

func (bot *Bot) Stop() error {
	err := bot.session.Close()
	if err != nil {
		return fmt.Errorf("error closing session: %w", err)
	}

	log.Println("Stopped bot")

	return nil
}

func (bot *Bot) AddCommand(command string, handler func(message *message.Message)) {
	bot.session.AddHandler(func(_ *discordgo.Session, msg *discordgo.MessageCreate) {
		if !strings.HasPrefix(msg.Content, commandPrefix+command) {
			return
		}

		handler(message.New(msg.Message, bot.session))
	})
}

func (bot *Bot) AddHandler(handler func(message *message.Message)) {
	bot.session.AddHandler(func(_ *discordgo.Session, msg *discordgo.MessageCreate) {
		handler(message.New(msg.Message, bot.session))
	})
}
