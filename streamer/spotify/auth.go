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

func NewAuthSpotify() *AuthenticatorSpotify {
	authSpotify := &AuthenticatorSpotify{State: "abc123"}
	authSpotify.Ch = make(chan *spotify.Client)
	return authSpotify
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

func (a *AuthenticatorSpotify) GetAuthURL() string {
	return a.Authenticator.AuthURL(a.State)
}

func (a *AuthenticatorSpotify) GetChannel() chan *spotify.Client {
	return a.Ch
}
