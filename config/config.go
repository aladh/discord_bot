package config

import (
	"fmt"
	"os"
)

type Config struct {
	DiscordToken        string
	SpotifyClientID     string
	SpotifyClientSecret string
	SpotifyRefreshToken string
	SpotifyPlaylistID   string
}

func FromEnv() (*Config, error) {
	discordToken, err := getEnvString("DISCORD_TOKEN")
	if err != nil {
		return nil, err
	}

	spotifyClientID, err := getEnvString("SPOTIFY_CLIENT_ID")
	if err != nil {
		return nil, err
	}

	spotifyClientSecret, err := getEnvString("SPOTIFY_CLIENT_SECRET")
	if err != nil {
		return nil, err
	}

	spotifyRefreshToken, err := getEnvString("SPOTIFY_REFRESH_TOKEN")
	if err != nil {
		return nil, err
	}

	spotifyPlaylistID, err := getEnvString("SPOTIFY_PLAYLIST_ID")
	if err != nil {
		return nil, err
	}

	return &Config{
		DiscordToken:        discordToken,
		SpotifyClientID:     spotifyClientID,
		SpotifyClientSecret: spotifyClientSecret,
		SpotifyRefreshToken: spotifyRefreshToken,
		SpotifyPlaylistID:   spotifyPlaylistID,
	}, nil
}

func getEnvString(name string) (string, error) {
	value, exists := os.LookupEnv(name)
	if !exists {
		return "", fmt.Errorf("environment variable %s not found", name)
	}

	return value, nil
}
