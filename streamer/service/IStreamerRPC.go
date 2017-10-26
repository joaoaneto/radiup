package service

import (
	"net/http"

	"github.com/joaoaneto/radiup/streamer/model"
	"github.com/joaoaneto/spotify"
	//"golang.org/x/oauth2"
)

/*RPC interfaces */
type ContentRPC interface {
	GetMusicData(client *spotify.Client, musicName string) ([]model.Music, error) //reformular
	GetPlaylistData(client *spotify.Client) (*spotify.SimplePlaylistPage, error)
}

type SocialRPC interface {
	//GetFollowers()
	GetInstant(client *spotify.Client) (model.Music, error)
	GetLastPlayedMusics(client *spotify.Client) ([]model.Music, error)
}

type AuthRPC interface {
	NewAuthenticator()
	GetAuthenticator() spotify.Authenticator
	NewClientAuth(w http.ResponseWriter, r *http.Request)
	GetAuthURL() string
	//GetChannel() chan *spotify.Client
}
