package service

import (
	"net/http"

	"github.com/joaoaneto/radiup/cycle"
	"github.com/joaoaneto/spotify"
	//"golang.org/x/oauth2"
)

/*RPC interfaces */
type ContentRPC interface {
	GetMusicData(client *spotify.Client, musicName string) ([]cycle.Music, error) //reformular
	GetPlaylistData(client *spotify.Client) (*spotify.SimplePlaylistPage, error)
}

type SocialRPC interface {
	//GetFollowers()
	GetInstant(client *spotify.Client) (cycle.Music, error)
	GetLastPlayedMusics(client *spotify.Client) ([]cycle.Music, error)
}

type AuthRPC interface {
	NewAuthenticator()
	GetAuthenticator() spotify.Authenticator
	NewClientAuth(w http.ResponseWriter, r *http.Request)
	GetAuthURL() string
	//GetChannel() chan *spotify.Client
}
