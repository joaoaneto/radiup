package spotify

import (
	"fmt"
	"log"
	"testing"

	"net/http"

	"github.com/joaoaneto/radiup/streamer"
	"github.com/zmb3/spotify"
)

func TestNewAuthenticator(t *testing.T) {
	i := AuthenticatorSpotify{}
	redirectURI := "http://localhost:8080/callback"
	i.NewAuthenticator(redirectURI)

	// Verify the error conditions

}

func TestSetAuthInfo(t *testing.T) {
	i := AuthenticatorSpotify{}
	redirectURI := "http://localhost:8080/callback"
	i.NewAuthenticator(redirectURI)
	auth := streamer.OAuthInfo{ClientID: "aaa", SecretKey: "000"}
	i.SetAuthInfo(auth)

	// Verify the error conditions

}

func TestNewClientAuth(t *testing.T) {
	i := AuthenticatorSpotify{}
	redirectURI := "http://localhost:8080/callback"
	i.NewAuthenticator(redirectURI)
	i.Ch = make(chan *spotify.Client)
	i.State = "abc123"

	http.HandleFunc("/callback", i.NewClientAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})

	go http.ListenAndServe(":8080", nil)

	url := i.Authenticator.AuthURL(i.State)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	client := <-i.Ch

	// Verify the error conditions

}
