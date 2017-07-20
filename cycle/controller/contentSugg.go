package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"time"

	"github.com/joaoaneto/radiup/cycle"
	"github.com/joaoaneto/radiup/cycle/repository/mongo"
	"github.com/joaoaneto/radiup/server"
)

func ShowContentSuggestionsHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("teasda")

	contentPersistor := mongo.NewPersistorContentSuggestion()
	fp1 := path.Join("templates", "base.html")
	fp2 := path.Join("templates", "content_info.html")
	t, err := template.ParseFiles(fp1, fp2)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, _ := contentPersistor.SearchAll(0)

	// send any data for template render through second argument in t.Execute()
	if err := t.Execute(w, result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegisterContentSuggestionsHandler(w http.ResponseWriter, r *http.Request) {

	contentPersistor := mongo.NewPersistorContentSuggestion()

	redirect := "/login"

	if r.Method == "GET" {

		sessionStore := server.GetSessionStore()
		session, _ := sessionStore.Store.Get(r, "cookie-name")

		auth, ok := session.Values["authenticated"].(bool)

		//verify if user is not authenticated, so redirect to login page
		if !ok || !auth {
			http.Redirect(w, r, redirect, 301)
			return
		}

		//contentPersistor := mongo.NewPersistorContentSuggestion()
		fp1 := path.Join("templates", "base.html")
		fp2 := path.Join("templates", "content_register.html")

		t, err := template.ParseFiles(fp1, fp2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := t.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	if r.Method == "POST" {

		r.ParseForm()

		//here... the idea is get the current user authenticated
		newUser := cycle.User{Name: "José Neto", Username: "netoax", Password: []byte("teste123"),
			BirthDay: time.Date(1995, time.August, 19, 0, 0, 0, 0, time.UTC), Email: "joao.alexandre@upe.br", Sex: "M"}

		newContentSuggestion := cycle.ContentSuggestion{Title: r.Form["title"][0], Description: r.Form["description"][0], ContentSuggestionUser: newUser,
			Votes: 0, Validated: false, Done: false}

		fmt.Println(newContentSuggestion)

		contentPersistor.Register(0, newContentSuggestion)

		http.Redirect(w, r, "/content/list", 301)
	}

}
