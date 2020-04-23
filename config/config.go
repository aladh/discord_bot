package config

import (
	"fmt"
	"os"
)

type Config struct {
	DiscordToken string
}

func New() (*Config, error) {
	discordToken, err := getEnvString("DISCORD_TOKEN")
	if err != nil {
		return nil, err
	}

	return &Config{DiscordToken: discordToken}, nil
}

func getEnvString(name string) (string, error) {
	value, exists := os.LookupEnv(name)
	if !exists {
		return "", fmt.Errorf("environment variable %s not found", name)
	}

	return value, nil
}
