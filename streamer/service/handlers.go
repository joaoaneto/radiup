package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joaoaneto/radiup/streamer/service/spotify"
)

type URI struct {
	Uri string `json:"uri"`
}

func GetMusicDataHandler(w http.ResponseWriter, r *http.Request) {

	var musicName = mux.Vars(r)["musicName"]
	//var nameUser = mux.Vars(r)["user"]

	spotifyStreamer := GetStreamerManager().Get("SPOTIFY")

	// search user according to nameUser var - AMQP
	// auth = NewAuthenticator()
	// generate client = auth.NewClient(nameUser)

	spotifyStreamer.ContentRPC.GetMusicData(nil, musicName)

}

func GetAuthURIHandler(w http.ResponseWriter, r *http.Request) {

	authSpotify := spotify.NewAuthSpotify()
	authSpotify.NewAuthenticator()

	uri := &URI{authSpotify.GetAuthURL()}

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)

	err := enc.Encode(uri)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(buf.Len()))
	w.WriteHeader(http.StatusOK)
	w.Write(buf.Bytes())

}
