package controller

import (
	"path"
	"net/http"
	"html/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	fp1 := path.Join("templates", "base.html")
	fp2 := path.Join("templates", "index.html")
	t, err := template.ParseFiles(fp1, fp2)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := t.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}