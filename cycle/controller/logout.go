package controller

import (
	"net/http"

	"github.com/joaoaneto/radiup/server"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	sessionStore := server.GetSessionStore()
	session, _ := sessionStore.Store.Get(r, "cookie-name")

	auth, ok := session.Values["authenticated"].(bool)

	if ok && auth {
		session.Values["authenticated"] = false
		session.Save(r, w)		
	}

}