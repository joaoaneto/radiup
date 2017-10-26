package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/oauth2"

	"github.com/gorilla/mux"
	"github.com/joaoaneto/radiup/streamer/service/spotify"
)

type URI struct {
	Uri string `json:"uri"`
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

func GetMusicDataHandler(w http.ResponseWriter, r *http.Request) {

	token := &oauth2.Token{}

	musicName := mux.Vars(r)["musicName"]
	userToken := r.Header.Get("Authorization")

	json.Unmarshal([]byte(userToken), token)
	tokenValues := strings.Split(userToken, " ")

	if len(tokenValues) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token.TokenType = tokenValues[0]
	token.AccessToken = tokenValues[1]

	spotifyStreamer := GetStreamerManager().Get("SPOTIFY")
	auth := spotifyStreamer.AuthRPC.GetAuthenticator()
	client := auth.NewClient(token)

	musicDataSpotify, err := spotifyStreamer.ContentRPC.GetMusicData(&client, musicName)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	musicData, _ := json.Marshal(musicDataSpotify)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(musicData)))
	w.WriteHeader(http.StatusOK)
	w.Write(musicData)

}

func GetPlaylistDataHandler(w http.ResponseWriter, r *http.Request) {

	token := &oauth2.Token{}

	tokenValues := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenValues) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token.TokenType = tokenValues[0]
	token.AccessToken = tokenValues[1]

	spotifyStreamer := GetStreamerManager().Get("SPOTIFY")
	auth := spotifyStreamer.AuthRPC.GetAuthenticator()
	client := auth.NewClient(token)

	playlistDataSpotify, err := spotifyStreamer.ContentRPC.GetPlaylistData(&client)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	playlistData, err := json.Marshal(playlistDataSpotify)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(playlistData)))
	w.WriteHeader(http.StatusOK)
	w.Write(playlistData)

}

func GetInstantMusicHandler(w http.ResponseWriter, r *http.Request) {

	token := &oauth2.Token{}

	tokenValues := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenValues) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token.TokenType = tokenValues[0]
	token.AccessToken = tokenValues[1]

	spotifyStreamer := GetStreamerManager().Get("SPOTIFY")
	auth := spotifyStreamer.AuthRPC.GetAuthenticator()
	client := auth.NewClient(token)

	instantMusicSpotify, err := spotifyStreamer.SocialRPC.GetInstant(&client)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	instantMusic, err := json.Marshal(instantMusicSpotify)
	if err != nil {
		w.WriteHeader(http.StatusProcessing)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(instantMusic)))
	w.WriteHeader(http.StatusOK)
	w.Write(instantMusic)

}

func GetLastPlayedMusicsHandler(w http.ResponseWriter, r *http.Request) {

	token := &oauth2.Token{}

	tokenValues := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenValues) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token.TokenType = tokenValues[0]
	token.AccessToken = tokenValues[1]

	spotifyStreamer := GetStreamerManager().Get("SPOTIFY")
	auth := spotifyStreamer.AuthRPC.GetAuthenticator()
	client := auth.NewClient(token)

	lastMusicsSpotify, _ := spotifyStreamer.SocialRPC.GetLastPlayedMusics(&client)

	lastPlayedMusics, _ := json.Marshal(lastMusicsSpotify)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(lastPlayedMusics)))
	w.WriteHeader(http.StatusOK)
	w.Write(lastPlayedMusics)

}
