package message

import (
	"fmt"
	"time"

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
		return nil, fmt.Errorf("error sending message: %w", err)
	}

	return New(msg, message.session), nil
}

func (message *Message) Edit(content string) error {
	_, err := message.session.ChannelMessageEdit(message.ChannelID, message.ID, content)
	if err != nil {
		return fmt.Errorf("error editing message: %w", err)
	}

	return nil
}

func (message *Message) React(reaction string) error {
	err := message.session.MessageReactionAdd(message.ChannelID, message.ID, reaction)
	if err != nil {
		return fmt.Errorf("error reacting to message: %w", err)
	}

	return nil
}

func (message *Message) CreatedAt() (time.Time, error) {
	timestamp, err := message.Timestamp.Parse()
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing timestamp: %w", err)
	}

	return timestamp, nil
}
