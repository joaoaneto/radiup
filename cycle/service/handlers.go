package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"github.com/joaoaneto/radiup/cycle"
	"github.com/joaoaneto/radiup/cycle/model"
	"github.com/joaoaneto/radiup/cycle/repository"
)

var Db *repository.MySQLConfig

func RegisterVoluntarySuggestionHandler()w http.ResponseWriter, r *http.Request){
	var vs cycle.VoluntarySuggestion{nil, nil, nil, nil}

	decoder := json.NewDecoder(r.Body)

	cp := repository.NewPersistorCycle()
	rc := cp.Search(0)
	if rc != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err := decoder.Decode(&vs)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	nvs := repository.NewPersistorVoluntarySuggestion()
	err = nvs.Register(rc.ID, vs)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)	
}

func GetVoluntarySuggestionHandler(w http.ResponseWriter, r *http.Request) {

	var localCycle cycle.Cycle{0, nil, nil, nil, nil, nil, nil, nil}

	cp := repository.NewPersistorCycle()	
	cp.Create(localCycle)

	listVS, err := cp.SearchAll(localCycle.ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(listVS)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}