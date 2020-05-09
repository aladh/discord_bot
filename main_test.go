package main

import (
	"log"
	"os"
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"

	"github.com/ali-l/discord_bot/config"
)

func TestSpotifyAddToPlaylist(t *testing.T) {
	_, enabled := os.LookupEnv("ENABLE_E2E_TESTS")
	if !enabled {
		t.Skip("End to end tests are not enabled")
	}

	cfg, err := config.FromEnv()
	if err != nil {
		t.Fatal(err)
	}

	sendAndDeleteSpotifyLink(t, cfg.DiscordToken, os.Getenv("DISCORD_CHANNEL_ID"), os.Getenv("SPOTIFY_TRACK_LINK"))

	spotifyClient := createSpotifyClient(cfg)

	track := getLastTrack(t, spotifyClient, cfg.SpotifyPlaylistID)
	expectedTrackName := os.Getenv("SPOTIFY_TRACK_NAME")

	if track.Name != expectedTrackName {
		t.Fatalf("track = %s, want %s", track.Name, expectedTrackName)
	}

	_, err = spotifyClient.RemoveTracksFromPlaylist(spotify.ID(cfg.SpotifyPlaylistID), track.ID)
	if err != nil {
		t.Fatalf("error removing track %s from playlist: %s", track.Name, err)
	}
}

func getLastTrack(t *testing.T, client spotify.Client, playlistID string) *spotify.FullTrack {
	playlist, err := client.GetPlaylist(spotify.ID(playlistID))
	if err != nil {
		t.Fatalf("error getting playlist: %s", err)
	}

	lastTrack := playlist.Tracks.Total - 1
	limit := int(1)

	tracks, err := client.GetPlaylistTracksOpt(spotify.ID(playlistID), &spotify.Options{Limit: &limit, Offset: &lastTrack}, "items(track(id,name))")
	if err != nil {
		t.Fatalf("error getting playlist tracks: %s", err)
	}

	return &tracks.Tracks[0].Track
}

func createSpotifyClient(cfg *config.Config) spotify.Client {
	auth := spotify.NewAuthenticator("", spotify.ScopePlaylistModifyPublic)
	auth.SetAuthInfo(cfg.SpotifyClientID, cfg.SpotifyClientSecret)
	client := auth.NewClient(&oauth2.Token{TokenType: "Bearer", RefreshToken: cfg.SpotifyRefreshToken})

	return client
}

func sendAndDeleteSpotifyLink(t *testing.T, token string, channelID string, trackLink string) {
	session := createDiscordSession(t, token)
	defer closeSession(session)

	msg, err := session.ChannelMessageSend(channelID, trackLink)
	if err != nil {
		t.Fatalf("error sending message: %s", err)
	}

	err = session.ChannelMessageDelete(channelID, msg.ID)
	if err != nil {
		t.Fatalf("error deleting message: %s", err)
	}
}

func createDiscordSession(t *testing.T, discordToken string) *discordgo.Session {
	session, err := discordgo.New("Bot " + discordToken)
	if err != nil {
		t.Fatalf("error initializng session: %s", err)
	}

	err = session.Open()
	if err != nil {
		t.Fatalf("error opening connection: %s", err)
	}

	return session
}

func closeSession(session *discordgo.Session) {
	err := session.Close()
	if err != nil {
		log.Fatalf("error closing session: %s\n", err)
	}
}
