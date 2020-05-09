package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ali-l/discord_bot/bot"
	"github.com/ali-l/discord_bot/config"
	"github.com/ali-l/discord_bot/ping"
	"github.com/ali-l/discord_bot/spotify"
)

func main() {
	cfg, err := config.FromEnv()
	if err != nil {
		log.Fatalln(err)
	}

	bt, err := bot.Start(cfg.DiscordToken)
	if err != nil {
		log.Fatalln(err)
	}
	defer bt.Stop()

	bt.AddCommand("ping", ping.ReplyWithLatency)
	bt.AddHandler(spotify.New(cfg.SpotifyClientID, cfg.SpotifyClientSecret, cfg.SpotifyRefreshToken, cfg.SpotifyPlaylistID).AddToPlaylist)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	<-stop
}
