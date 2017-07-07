package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"path"
	"time"

	"github.com/joaoaneto/radiup/cycle"
	"github.com/joaoaneto/radiup/cycle/repository/mongo"
)

func ShowContentSuggestionsHandler(w http.ResponseWriter, r *http.Request) {

	contentPersistor := mongo.NewPersistorContentSuggestion()
	fp1 := path.Join("templates", "base.html")
	fp2 := path.Join("templates", "content_info.html")
	t, err := template.ParseFiles(fp1, fp2)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, _ := contentPersistor.SearchAll()

	// send any data for template render through second argument in t.Execute()
	if err := t.Execute(w, result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegisterContentSuggestionsHandler(w http.ResponseWriter, r *http.Request) {

	//contentPersistor := mongo.NewPersistorContentSuggestion()
	fp1 := path.Join("templates", "base.html")
	fp2 := path.Join("templates", "content_register.html")

	contentPersistor := mongo.NewPersistorContentSuggestion()

	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		t, err := template.ParseFiles(fp1, fp2)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := t.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		r.ParseForm()
		newUser := cycle.User{Name: "José Neto", Username: "netoax", Password: []byte("teste123"),
			BirthDay: time.Date(1995, time.August, 19, 0, 0, 0, 0, time.UTC), Email: "joao.alexandre@upe.br", Sex: 'M'}

		newContentSuggestion := cycle.ContentSuggestion{Title: r.Form["title"][0], Description: r.Form["description"][0], ContentSuggestionUser: newUser,
			Votes: 0, Validated: false, Done: false}

		fmt.Println(newContentSuggestion)

		contentPersistor.Register(newContentSuggestion)

		http.Redirect(w, r, "/content/list", 301)
	}

}
