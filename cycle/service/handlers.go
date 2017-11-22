package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/joaoaneto/radiup/cycle/model"
	"github.com/joaoaneto/radiup/cycle/repository/mongo"
	"github.com/joaoaneto/radiup/cycle/service/spotify"
	"github.com/joaoaneto/radiup/user/repository"
)

var Db *repository.MySQLConfig

func RegisterVoluntarySuggestionHandler(w http.ResponseWriter, r *http.Request) {

	music := model.Music{PlayedAt: time.Now()}
	voluntarySuggestion := model.VoluntarySuggestion{Track: music, Timestamp: time.Now()}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&voluntarySuggestion)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nvs := mongo.NewPersistorVoluntarySuggestion()

	daeler := spotify.NewVoluntarySuggestionDealer()

	userHasVoted, err := daeler.VerifyUserVote(0, voluntarySuggestion.Track.ID, voluntarySuggestion.Users[0])
	if userHasVoted && err == nil { //if user already choose or voted
		log.Println("User already choose this music.")
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if !userHasVoted && err == nil { //if suggestion exists but user don't make it
		log.Println("There is suggestion, but user don't make it.")
		sugg, _ := nvs.Search(0, voluntarySuggestion.Track.ID)
		sugg.Users = append(sugg.Users, voluntarySuggestion.Users[0])
		sugg.Votes++
		nvs.Update(sugg)
		w.WriteHeader(http.StatusOK)
		return
	}

	err = nvs.Register(0, voluntarySuggestion)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetVoluntarySuggestionHandler(w http.ResponseWriter, r *http.Request) {

	vsp := mongo.NewPersistorVoluntarySuggestion()

	voluntarySuggestions, err := vsp.SearchAll(0)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(voluntarySuggestions)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func RegisterVoteHandler(w http.ResponseWriter, r *http.Request) {

	username := r.URL.Query().Get("username")
	musicId := r.URL.Query().Get("musicId")

	pvs := mongo.NewPersistorVoluntarySuggestion()

	voluntarySuggestion, err := pvs.Search(0, musicId)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	voluntarySuggestion.Votes++
	voluntarySuggestion.Users = append(voluntarySuggestion.Users, username)
	pvs.Update(voluntarySuggestion)

	log.Println(string(musicId) + " " + string(username))
	w.WriteHeader(http.StatusOK)
}
