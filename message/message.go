package message

import (
	"github.com/bwmarrin/discordgo"
)

type Message struct {
	*discordgo.Message
	client Client
}

type Client interface {
	SendMessage(channelID string, content string) (*Message, error)
	EditMessage(channelID string, messageID string, content string) error
}

func New(message *discordgo.Message, client Client) *Message {
	return &Message{
		Message: message,
		client:  client,
	}
}

func (message *Message) Reply(content string) (*Message, error) {
	msg, err := message.client.SendMessage(message.ChannelID, content)
	if err != nil {
		return nil, err
	}

	return msg, nil
}

func (message *Message) Edit(content string) error {
	err := message.client.EditMessage(message.ChannelID, message.ID, content)
	if err != nil {
		return err
	}

	return nil
}
