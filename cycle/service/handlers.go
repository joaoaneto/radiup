package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/joaoaneto/radiup/user/model"
	"github.com/joaoaneto/radiup/user/repository"
)

/*
curl -H "Content-Type: application/json" -X POST -d '{"Name":"João Neto", "Username": "netoax",
"Password":"teste", "BirthDay":"2015-09-15T14:00:12-00:00", "Email":"joao.alexandre@upe.br",
"Sex":"Masculino", "SpotifyToken": {"AccessToken": "token-test", "TokenType": "aloka",
"RefreshToken": "cê é louco", "Expiry": "2015-09-15T14:00:12-00:00"}}' http://localhost:6767/register
*/

var Db *repository.MySQLConfig

//var MessagingClient messaging.IMessagingClient

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

	user := model.User{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println(user)

	sup := repository.NewSimpleUserPersistor(Db)

	err = sup.Create(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

/*func notifyTest(account model.User) {
	go func(account model.User) {
		notification := model.
	}
}*/

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	// Get user in db
	var userId = mux.Vars(r)["userId"]

	sup := repository.NewSimpleUserPersistor(Db)

	user, err := sup.Search(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)

}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {

	userId := mux.Vars(r)["userId"]
	sup := repository.NewSimpleUserPersistor(Db)

	user, err := sup.Search(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = sup.Update(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	var userId = mux.Vars(r)["userId"]

	sup := repository.NewSimpleUserPersistor(Db)

	err := sup.Remove(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}

func Authenticate(w http.ResponseWriter, r *http.Request) {

	test := model.Teste{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&test)
	if err != nil {
		return
	}

	fmt.Println(test)

}
