package ping

import (
	"log"

	"github.com/ali-l/discord_bot_go/message"
)

func Handle(message *message.Message) {
	_, err := message.Reply("Pong!")
	if err != nil {
		log.Println(err)
	}
}
