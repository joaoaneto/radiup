package streamer

import (
	"github.com/zmb3/spotify"
	"net/http"
)

/*RPC interfaces */
type ContentRPC interface {
	GetMusicData()
	GetPlaylistData()
}

type SocialRPC interface {
	GetFollowers()
	GetInstant()
}

type AuthRPC interface {
	NewAuthenticator(redirectURI string)
	SetAuthInfo(auth OAuthInfo)
	NewClientAuth(w http.ResponseWriter, r *http.Request)
	GetAuthURL() string
	GetChannel() chan *spotify.Client
}
