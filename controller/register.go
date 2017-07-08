package controller

import (
	"html/template"
	"net/http"
	"path"
	"github.com/joaoaneto/radiup/streamer"
	"fmt"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	spotifyStreamer := streamer.GetStreamerManager().Get("SPOTIFY")
	spotifyStreamer.AuthRPC.NewAuthenticator("http://localhost:8080/callback")
	oAuthTest := streamer.OAuthInfo{ClientID: "42d13a4cacae480189b2702e48d6879a", SecretKey: "f0864a30cca443c4b33b940940285d87"}
	spotifyStreamer.AuthRPC.SetAuthInfo(oAuthTest)
	
	fp1 := path.Join("templates", "base.html")
	fp2 := path.Join("templates", "register_form.html")
	t, err := template.ParseFiles(fp1, fp2)

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
