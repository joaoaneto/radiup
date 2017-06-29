package streamer

import (
	"net/http"

	"github.com/joaoaneto/radiup/cycle"
	"github.com/zmb3/spotify"
)

/*RPC interfaces */
type ContentRPC interface {
	GetMusicData()
	GetPlaylistData()
}

type SocialRPC interface {
	//GetFollowers()
	GetInstant(client *spotify.Client) (cycle.Music, error)
}

type AuthRPC interface {
	NewAuthenticator(redirectURI string)
	SetAuthInfo(auth OAuthInfo)
	NewClientAuth(w http.ResponseWriter, r *http.Request)
	GetAuthURL() string
	GetChannel() chan *spotify.Client
}
