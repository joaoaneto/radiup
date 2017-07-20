package main

import (
	//"fmt"

	"net/http"

	"github.com/joaoaneto/radiup/cycle/controller"
	"github.com/joaoaneto/radiup/streamer"
	"github.com/joaoaneto/radiup/streamer/spotify"
)

func main() {

	// Improve this, please.
	// I don't found better place for calling this function.
	// By: JN
	spotify.Start()

	spotifyStreamer := streamer.GetStreamerManager().Get("SPOTIFY")

	spotifyStreamer.AuthRPC.NewAuthenticator("http://localhost:8080/register")
	oAuthTest := streamer.OAuthInfo{ClientID: "42d13a4cacae480189b2702e48d6879a", SecretKey: "f0864a30cca443c4b33b940940285d87"}
	spotifyStreamer.AuthRPC.SetAuthInfo(oAuthTest)

	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/callback", spotifyStreamer.AuthRPC.NewClientAuth)
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/register", controller.RegisterHandler)
	http.HandleFunc("/content/list", controller.ShowContentSuggestionsHandler)
	http.HandleFunc("/content/register", controller.RegisterContentSuggestionsHandler)
	http.HandleFunc("/logout", controller.LogoutHandler)
	http.HandleFunc("/voluntary/list", controller.ShowVoluntarySuggestionsHandler)
	http.HandleFunc("/voluntary/register", controller.RegisterVoluntarySuggestionsHandler)
	http.HandleFunc("/voluntary/search", controller.SearchVoluntarySuggestionsHandler)
	http.ListenAndServe(":8080", nil)

}
