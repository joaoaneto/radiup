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

func GetVoluntarySuggestionsHandler(w http.ResponseWriter, r *http.Request) {

	/* Get user in db
	var userId = mux.Vars(r)["userId"]

	sup := repository.NewSimpleUserPersistor(Db)
	*/
	var localCycle cycle.Cycle{0, nil, nil, nil, nil, nil, nil, nil}

	cp := repository.NewCyclePersistor(Db)	
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