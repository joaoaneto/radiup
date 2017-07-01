package controller

import (
	"html/template"
	"net/http"
	"path"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	fp1 := path.Join("templates", "base.html")
	fp2 := path.Join("templates", "login_form.html")
	t, err := template.ParseFiles(fp1, fp2)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// send any data for template render through second argument in t.Execute()
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
