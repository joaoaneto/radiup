package controller

import (
	"fmt"
	"net/http"
	"path"
	"html/template"
	"time"

	"github.com/joaoaneto/radiup/cycle"
	"github.com/joaoaneto/radiup/cycle/business"
	"github.com/joaoaneto/radiup/cycle/repository/mongo"
	"github.com/joaoaneto/radiup/streamer"
	"github.com/joaoaneto/spotify"
)

func ShowVoluntarySuggestionsHandler(w http.ResponseWriter, r *http.Request) {

	fp1 := path.Join("templates", "base.html")
	fp2 := path.Join("templates", "voluntary_list.html")
	t, err := template.ParseFiles(fp1, fp2)

	spotifyStreamer := streamer.GetStreamerManager().Get("SPOTIFY")
	auth := spotifyStreamer.AuthRPC.GetAuthenticator()

	userPersistor := mongo.NewPersistorSimpleUser()
	musicPersistor := mongo.NewPersistorMusic()
	persistorVSugg := mongo.NewPersistorVoluntarySuggestion()
		
	user, _ := userPersistor.Search("netoax")
	business.CheckTokenExpiry(user)
	client := auth.NewClient(user.AuthSpotify)

	var listVSugg []cycle.VoluntarySuggestion

	if r.Method == "POST" {

		var music cycle.Music
		var id spotify.ID
		r.ParseForm()

		value := r.Form["number"][0]

		id = spotify.ID(value)

		fmt.Println(id)

		track, _ := client.GetTrack(id)
		
		var artistsList []string

	    music.ID = track.SimpleTrack.ID.String()

	    music.Name = track.SimpleTrack.Name

	    for _, a := range track.SimpleTrack.Artists {
	      artistsList = append(artistsList, a.Name)
	    }

	    music.Artist = artistsList

		musicPersistor.Register(music)
		fmt.Println("opa")
		voluntarySuggestion := cycle.VoluntarySuggestion{music, user.SimpleUser, 0, time.Now()}
		persistorVSugg.Register(0, voluntarySuggestion)
		fmt.Println("Sugestão adicionada")
	}

	listVSugg, _ = persistorVSugg.SearchAll(0)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, listVSugg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func SearchVoluntarySuggestionsHandler(w http.ResponseWriter, r *http.Request) {

	fp1 := path.Join("templates", "base.html")
	fp2 := path.Join("templates", "voluntary_search.html")
	t, err := template.ParseFiles(fp1, fp2)

	spotifyStreamer := streamer.GetStreamerManager().Get("SPOTIFY")
	auth := spotifyStreamer.AuthRPC.GetAuthenticator()
	userPersistor := mongo.NewPersistorSimpleUser()

	var search []cycle.Music

	if r.Method == "POST" {
		r.ParseForm()

		user, _ := userPersistor.Search("netoax")
		business.CheckTokenExpiry(user)
		searchMusic := r.Form["search"][0]
		fmt.Println("Música procurada: ", searchMusic)
		client := auth.NewClient(user.AuthSpotify)
		search, _ = spotifyStreamer.ContentRPC.GetMusicData(&client, searchMusic)

	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, search); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}



func RegisterVoluntarySuggestionsHandler(w http.ResponseWriter, r *http.Request) {

	fp1 := path.Join("templates", "base.html")
	fp2 := path.Join("templates", "voluntary_register.html")
	t, err := template.ParseFiles(fp1, fp2)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}