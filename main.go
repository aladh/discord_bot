package main

import (
	"log"

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

	bt, err := bot.New(cfg.DiscordToken)
	if err != nil {
		log.Fatalln(err)
	}

	bt.AddCommand("ping", ping.ReplyWithLatency)
	bt.AddHandler(spotify.New(cfg.SpotifyClientID, cfg.SpotifyClientSecret, cfg.SpotifyRefreshToken, cfg.SpotifyPlaylistID).AddToPlaylist)

	err = bt.Start()
	if err != nil {
		log.Fatalf("error starting bot: %s\n", err)
	}
}
