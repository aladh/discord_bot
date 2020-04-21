package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

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

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-stop
}
