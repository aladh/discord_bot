package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
	DiscordToken string
}

func New() (*Config, error) {
	discordToken, err := getEnvString("DISCORD_TOKEN")
	if err != nil {
		return nil, fmt.Errorf("failed to generate config: %w", err)
	}

	return &Config{
		DiscordToken: discordToken,
	}, nil
}

func getEnvString(name string) (string, error) {
	value, exists := os.LookupEnv(name)
	if !exists {
		return "", errors.New(fmt.Sprintf("environment variable %s not found", name))
	}

	return value, nil
}
