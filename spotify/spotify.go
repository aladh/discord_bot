package spotify

import (
	"context"
	"fmt"
	"log"
	"regexp"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"

	"github.com/aladh/discord_bot/message"
)

const trackIDPattern = "https://open\\.spotify\\.com/track/([0-9A-Za-z]*)"

type Client struct {
	*spotify.Client
	playlistID   spotify.ID
	trackIDRegex *regexp.Regexp
}

func New(clientID, clientSecret, refreshToken, playlistID string) *Client {
	auth := spotifyauth.New(
		spotifyauth.WithClientID(clientID),
		spotifyauth.WithClientSecret(clientSecret),
		spotifyauth.WithScopes(spotifyauth.ScopePlaylistModifyPublic),
	)

	httpClient := auth.Client(context.Background(), &oauth2.Token{TokenType: "Bearer", RefreshToken: refreshToken})
	client := spotify.New(httpClient)

	return &Client{Client: client, playlistID: spotify.ID(playlistID), trackIDRegex: regexp.MustCompile(trackIDPattern)}
}

func (client *Client) AddToPlaylist(message *message.Message) {
	trackID, err := extractTrackID(client.trackIDRegex, message.Content)
	if err != nil {
		return
	}

	_, err = client.AddTracksToPlaylist(context.Background(), client.playlistID, trackID)
	if err != nil {
		log.Println(err)
	}
}

func extractTrackID(trackIDRegex *regexp.Regexp, trackURL string) (spotify.ID, error) {
	matches := trackIDRegex.FindStringSubmatch(trackURL)
	numMatches := len(matches)

	if numMatches > 0 {
		return spotify.ID(matches[numMatches-1]), nil
	}

	return "", fmt.Errorf("track ID not found")
}
