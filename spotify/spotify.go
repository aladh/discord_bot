package spotify

import (
	"fmt"
	"log"
	"regexp"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"

	"github.com/ali-l/discord_bot_go/message"
)

const trackIDPattern = "open.spotify.com/track/([0-9A-Za-z_-]*)"

type Client struct {
	spotify.Client
	playlistID   string
	trackIDRegex *regexp.Regexp
}

func New(clientID, clientSecret, refreshToken, playlistID string) *Client {
	token := &oauth2.Token{
		TokenType:    "Bearer",
		RefreshToken: refreshToken,
	}

	auth := spotify.NewAuthenticator("", spotify.ScopePlaylistModifyPublic)

	auth.SetAuthInfo(clientID, clientSecret)

	client := auth.NewClient(token)

	return &Client{Client: client, playlistID: playlistID, trackIDRegex: regexp.MustCompile(trackIDPattern)}
}

func (client *Client) PlaylistAdder(message *message.Message) {
	trackID, err := extractTrackID(client.trackIDRegex, message.Content)
	if err != nil {
		return
	}

	_, err = client.AddTracksToPlaylist(spotify.ID(client.playlistID), trackID)
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
