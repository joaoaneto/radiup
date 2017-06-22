package spotify

import (
	"net/http"

	"github.com/joaoaneto/radiup/streamer"
	"github.com/zmb3/spotify"
)

// AuthenticatorSpotify ...
type AuthenticatorSpotify struct {
	Authenticator spotify.Authenticator
	Ch            (chan *spotify.Client)
	State         string
}

// NewAuthenticator ...
func (a *AuthenticatorSpotify) NewAuthenticator(
	redirectURI string) {

	a.Authenticator = spotify.NewAuthenticator(redirectURI,
		spotify.ScopeUserReadCurrentlyPlaying)

}

// SetAuthInfo ...
func (a *AuthenticatorSpotify) SetAuthInfo(auth streamer.OAuthInfo) {

	a.Authenticator.SetAuthInfo(auth.ClientID, auth.SecretKey)

}

// NewClientAuth ...
func (a *AuthenticatorSpotify) NewClientAuth(w http.ResponseWriter, r *http.Request) {

	tok, _ := a.Authenticator.Token(a.State, r)
	client := a.Authenticator.NewClient(tok)

	a.Ch <- &client

}
