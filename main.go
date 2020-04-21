package main

import (
	"log"

	"github.com/ali-l/discord_bot_go/bot"
	"github.com/ali-l/discord_bot_go/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err.Error())
	}

	bt, err := bot.Start(cfg.DiscordToken)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer bt.Stop()
}
