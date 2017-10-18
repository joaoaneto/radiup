package service

import (
	"log"
	"net/http"

	"github.com/joaoaneto/radiup/streamer/service/spotify"
)

func StartServer(port string) {

	authSpotify := spotify.NewAuthSpotify()
	authSpotify.NewAuthenticator()
	authSpotify.CallbackFunc = authSpotify.NewClientAuth

	r := NewRouter()
	http.Handle("/", r)
	http.Handle("/callback", authSpotify.CallbackFunc)

	log.Println("Start HTTP Server at " + port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Println("An error ocurred starting HTTP listener at port: " + port)
		log.Println("Error: " + err.Error())
	}

}
