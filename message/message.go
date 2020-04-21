package message

import (
	"github.com/bwmarrin/discordgo"
)

type Message struct {
	*discordgo.Message
	client Client
}

type Client interface {
	SendMessage(string, string) (*Message, error)
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
