package ping

import (
	"fmt"
	"log"

	"github.com/aladh/discord_bot/message"
)

func ReplyWithLatency(message *message.Message) {
	reply, err := message.Reply("Pong!")
	if err != nil {
		log.Println(err)
		return
	}

	messageCreatedAt, err := message.CreatedAt()
	if err != nil {
		log.Println(err)
		return
	}

	replyCreatedAt, err := reply.CreatedAt()
	if err != nil {
		log.Println(err)
		return
	}

	err = reply.Edit(fmt.Sprintf("Pong! (%dms)", replyCreatedAt.Sub(messageCreatedAt).Milliseconds()))
	if err != nil {
		log.Println(err)
	}
}
