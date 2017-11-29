package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	//"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

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

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {

	userType := r.Header.Get("user-type")

	if(userType == "1"){
		
		simpleUser := model.SimpleUser{/*BirthDay: time.Now()*/}

		decoder := json.NewDecoder(r.Body)
		
		err := decoder.Decode(&simpleUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println(r.Body)

		sup := repository.NewSimpleUserPersistor(Db)

		err = sup.Create(&simpleUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		
		adminUser := model.AdminUser{/*BidrthDay: time.Now()*/}

		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&adminUser)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		fmt.Println(r.Body)

		aup := repository.NewAdminUserPersistor(Db)

		err = aup.Create(&adminUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
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

func Authentication(w http.ResponseWriter, r *http.Request) {

	// TO-DO ----> Generate JWT

	var auth Auth

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&auth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("Lala: ", auth)

	authStatus := userAuthenticate(auth.Username, auth.Password)

	data, _ := json.Marshal(authStatus)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func generateHash(password string) []byte {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	return hash

}

func userAuthenticate(login string, password string) bool {

	simpleUserPersistor := repository.NewSimpleUserPersistor(Db)
	simpleUser, err := simpleUserPersistor.Search(login)
	if err != nil {
		return false
	}

	/*if err := bcrypt.CompareHashAndPassword(simpleUser.SimpleUser.Password, []byte(password)); err != nil {
		return false
	}*/

	fmt.Println(simpleUser)

	if simpleUser.Password == password {
		return true
	}

	return false
}
