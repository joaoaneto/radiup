package service

import (
	"net/http"
	"strings"

	cycleMongo "github.com/joaoaneto/radiup/cycle/repository/mongo"
	streamer "github.com/joaoaneto/radiup/streamer/service"
	"github.com/joaoaneto/spotify"
	"golang.org/x/oauth2"
)

func GenerateVolSuggestionsPlaylistHandler(w http.ResponseWriter, r *http.Request) {

	pvs := cycleMongo.NewPersistorVoluntarySuggestion()
	voluntarySuggestions, _ := pvs.SearchAll(0)

	token := &oauth2.Token{}

	tokenValues := strings.Split(r.Header.Get("Authorization"), " ")
	if len(tokenValues) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	token.TokenType = tokenValues[0]
	token.AccessToken = tokenValues[1]

	spotifyStreamer := streamer.GetStreamerManager().Get("SPOTIFY")
	auth := spotifyStreamer.AuthRPC.GetAuthenticator()
	client := auth.NewClient(token)

	user, _ := client.CurrentUser()

	playlist, _ := client.CreatePlaylistForUser(user.ID, "RadiUP", true)

	for _, v := range voluntarySuggestions {
		client.AddTracksToPlaylist(user.ID, playlist.ID, spotify.ID(v.Track.ID))
	}
}
