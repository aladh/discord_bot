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

	err = reply.Edit(fmt.Sprintf("Pong! (%dms)", reply.Timestamp.Sub(message.Timestamp).Milliseconds()))
	if err != nil {
		log.Println(err)
	}
}
