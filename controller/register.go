package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/joaoaneto/radiup/cycle"
	"github.com/joaoaneto/radiup/cycle/repository/mongo"
	"github.com/joaoaneto/radiup/streamer"
	"golang.org/x/crypto/bcrypt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	spotifyStreamer := streamer.GetStreamerManager().Get("SPOTIFY")
	spotifyStreamer.AuthRPC.NewAuthenticator("http://localhost:8080/callback")
	oAuthTest := streamer.OAuthInfo{ClientID: "42d13a4cacae480189b2702e48d6879a", SecretKey: "f0864a30cca443c4b33b940940285d87"}
	spotifyStreamer.AuthRPC.SetAuthInfo(oAuthTest)

	fp1 := path.Join("templates", "base.html")
	fp2 := path.Join("templates", "register_form.html")
	t, err := template.ParseFiles(fp1, fp2)

	userPersistor := mongo.NewPersistorUser()

	if r.Method == "POST" {
		r.ParseForm()

		//get data from template form
		name := r.Form["name"][0]
		email := r.Form["email"][0]
		birthDate, _ := time.Parse("2006-01-02", r.Form["date"][0])
		username := r.Form["username"][0]
		password, _ := authConfirmPassword(r.Form["password"][0], r.Form["passwordConfirm"][0])
		sexo := r.Form["optionsRadios"][0]

		fmt.Println(name, email, t, username, password, sexo)

		user := cycle.User{name, username, password, birthDate, email, sexo}

		userPersistor.Create(user)

		fmt.Println(user)

		http.Redirect(w, r, "/login", 301)

	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get URL
	templateUrl := spotifyStreamer.AuthRPC.GetAuthURL()

	// send any data for template render through second argument in t.Execute()
	if err := t.Execute(w, templateUrl); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func authConfirmPassword(pwd string, pwdConn string) ([]byte, bool) {

	if pwd == pwdConn {
		hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		return hash, true
	}

	return nil, false

}
