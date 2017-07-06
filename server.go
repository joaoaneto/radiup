package main

import (
	"net/http"

	"github.com/joaoaneto/radiup/controller"
	"github.com/joaoaneto/radiup/streamer"
	"github.com/joaoaneto/radiup/streamer/spotify"
)

func RunServer() {

	spotify.Start()
	spotifyStreamer := streamer.GetStreamerManager().Get("SPOTIFY")

	http.HandleFunc("/callback", spotifyStreamer.AuthRPC.NewClientAuth)
	http.HandleFunc("/", controller.LoginHandler)
	http.HandleFunc("/register", controller.RegisterHandler)

	http.ListenAndServe(":8080", nil)

}