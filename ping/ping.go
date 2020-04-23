package ping

import (
	"fmt"
	"log"
	"time"

	"github.com/ali-l/discord_bot_go/message"
)

func Handler(message *message.Message) {
	reply, err := message.Reply("Pong!")
	if err != nil {
		log.Println(err)
		return
	}

	timestamp, err := reply.Timestamp.Parse()
	if err != nil {
		log.Printf("error parsing timestamp: %s\n", err)
		return
	}

	err = reply.Edit(fmt.Sprintf("Pong! (%dms)", time.Since(timestamp).Milliseconds()))
	if err != nil {
		log.Println(err)
	}
}
