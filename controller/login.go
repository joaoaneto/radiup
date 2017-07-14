package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/joaoaneto/radiup/cycle/repository/mongo"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	fp1 := path.Join("templates", "base.html")
	fp2 := path.Join("templates", "login_form.html")
	t, err := template.ParseFiles(fp1, fp2)

	if r.Method == "POST" {
		r.ParseForm()
		login := r.Form["username"][0]
		password := r.Form["password"][0]
		if userAuthenticate(login, password) {
			fmt.Println("Ok")
			http.Redirect(w, r, "/content/list", 301)
		} else {
			fmt.Println("Wrong")
		}
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send any data for template render through second argument in t.Execute()
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func generateHash(password string) []byte {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatal(err)
	}

	return hash

}

func userAuthenticate(login string, password string) bool {

	simpleUserPersistor := mongo.NewPersistorSimpleUser()
	simpleUser, err := simpleUserPersistor.Search(login)
	if err != nil {
		return false
	}

	if err := bcrypt.CompareHashAndPassword(simpleUser.SimpleUser.Password, []byte(password)); err != nil {
		return false
	}

	return true

}
