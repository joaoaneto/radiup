package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/joaoaneto/radiup/cycle/model"
	"github.com/joaoaneto/radiup/cycle/repository/mongo"
	"github.com/joaoaneto/radiup/user/repository"
)

var Db *repository.MySQLConfig

func RegisterVoluntarySuggestionHandler(w http.ResponseWriter, r *http.Request) {

	voluntarySuggestion := model.VoluntarySuggestion{}

	decoder := json.NewDecoder(r.Body)

	//cp := mongo.NewPersistorCycle()

	/*cycle, err := cp.Search(0)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}*/

	err := decoder.Decode(&voluntarySuggestion)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nvs := mongo.NewPersistorVoluntarySuggestion()

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
