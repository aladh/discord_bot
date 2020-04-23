package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ali-l/discord_bot_go/bot"
	"github.com/ali-l/discord_bot_go/config"
	"github.com/ali-l/discord_bot_go/ping"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	bt, err := bot.Start(cfg.DiscordToken)
	if err != nil {
		log.Fatalln(err)
	}
	defer bt.Stop()

	bt.AddCommand("ping", ping.Handle)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-stop
}
