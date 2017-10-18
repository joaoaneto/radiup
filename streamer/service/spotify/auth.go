package spotify

import (
	"net/http"

	"github.com/joaoaneto/radiup/streamer/model"
	"github.com/joaoaneto/spotify"
	//"golang.org/x/oauth2"
)

var redirectURI = "http://localhost:6868/callback"

var spotifyOAuth = model.OAuthInfo{
	ClientID:  "bee86ebf12534e3eb4aca20742533c81",
	SecretKey: "952210d31c2c49d09a0c1ac163cc8116",
}

// AuthenticatorSpotify ...
type AuthenticatorSpotify struct {
	Authenticator spotify.Authenticator
	Client        spotify.Client
	State         string
	CallbackFunc  http.HandlerFunc
}

func NewAuthSpotify() *AuthenticatorSpotify {
	authSpotify := &AuthenticatorSpotify{State: "abc123"}
	//authSpotify.Ch = make(chan *spotify.Client)[]
	return authSpotify
}

// NewAuthenticator ...
func (a *AuthenticatorSpotify) NewAuthenticator() {

	a.Authenticator = spotify.NewAuthenticator(redirectURI,
		spotify.ScopeUserReadCurrentlyPlaying,
		spotify.ScopePlaylistReadCollaborative,
		spotify.ScopePlaylistReadPrivate,
		spotify.ScopePlaylistModifyPrivate,
		spotify.ScopePlaylistModifyPublic,
		spotify.ScopeUserModifyPlaybackState,
		spotify.ScopeUserFollowRead,
		spotify.ScopeUserLibraryRead,
		spotify.ScopeUserReadBirthdate,
		spotify.ScopeUserReadEmail,
		spotify.ScopeUserReadPlaybackState,
		spotify.ScopeUserReadPrivate,
		spotify.ScopeUserReadRecentlyPlayed)

	a.Authenticator.SetAuthInfo(spotifyOAuth.ClientID, spotifyOAuth.SecretKey)
}

func (a *AuthenticatorSpotify) GetAuthenticator() spotify.Authenticator {
	return a.Authenticator
}

func (a *AuthenticatorSpotify) NewClientAuth(w http.ResponseWriter, r *http.Request) {
	tok, _ := a.Authenticator.Token(a.State, r)
	a.Client = a.Authenticator.NewClient(tok)

	//notify user - OAuth2 Token
}

func (a *AuthenticatorSpotify) GetAuthURL() string {
	return a.Authenticator.AuthURL(a.State)
}

/*func (a *AuthenticatorSpotify) GetChannel() chan *spotify.Client {
	return a.Ch
}*/
