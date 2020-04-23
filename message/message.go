package message

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type Message struct {
	*discordgo.Message
	session *discordgo.Session
}

func New(message *discordgo.Message, session *discordgo.Session) *Message {
	return &Message{Message: message, session: session}
}

func (message *Message) Reply(content string) (*Message, error) {
	msg, err := message.session.ChannelMessageSend(message.ChannelID, content)
	if err != nil {
		return nil, fmt.Errorf("failed to send message: %w", err)
	}

	return New(msg, message.session), nil
}

func (message *Message) Edit(content string) error {
	_, err := message.session.ChannelMessageEdit(message.ChannelID, message.ID, content)
	if err != nil {
		return fmt.Errorf("failed to edit message: %w", err)
	}

	return nil
}
