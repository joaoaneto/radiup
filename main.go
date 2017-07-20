package main

import (
	"fmt"

	"net/http"
	//"github.com/joaoaneto/radiup/cycle"
	"github.com/joaoaneto/radiup/cycle/controller"
	//"github.com/joaoaneto/radiup/cycle/repository/mongo"
	cycleBusiness "github.com/joaoaneto/radiup/cycle/business"
	"github.com/joaoaneto/radiup/streamer"
	"github.com/joaoaneto/radiup/streamer/spotify"
	"github.com/joaoaneto/radiup/playlist"
)

func main() {

	// Improve this, please.
	// I don't found better place for calling this function.
	// By: JN
	spotify.Start()

	spotifyStreamer := streamer.GetStreamerManager().Get("SPOTIFY")

	spotifyStreamer.AuthRPC.NewAuthenticator("http://localhost:8080/register")
	oAuthTest := streamer.OAuthInfo{ClientID: "42d13a4cacae480189b2702e48d6879a", SecretKey: "f0864a30cca443c4b33b940940285d87"}
	spotifyStreamer.AuthRPC.SetAuthInfo(oAuthTest)

	http.HandleFunc("/", controller.IndexHandler)
	http.HandleFunc("/callback", spotifyStreamer.AuthRPC.NewClientAuth)
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/register", controller.RegisterHandler)
	http.HandleFunc("/content/list", controller.ShowContentSuggestionsHandler)
	http.HandleFunc("/content/register", controller.RegisterContentSuggestionsHandler)
	http.HandleFunc("/logout", controller.LogoutHandler)
	http.HandleFunc("/voluntary/list", controller.ShowVoluntarySuggestionsHandler)
	http.HandleFunc("/voluntary/register", controller.RegisterVoluntarySuggestionsHandler)
	http.HandleFunc("/voluntary/search", controller.SearchVoluntarySuggestionsHandler)
	http.ListenAndServe(":8080", nil)

	/*url := spotifyStreamer.AuthRPC.GetAuthURL()

	fmt.Println("Please, use this url for auth:")
	fmt.Println(url)

	client := <-spotifyStreamer.AuthRPC.GetChannel()
	tkn, _ := client.Token()


	usr := cycle.User{Username: "radiupapp"}
	adm := cycle.AdminUser{AdminUser: usr, AuthSpotify: tkn}

	admRep := mongo.NewPersistorAdminUser()
	admRep.Create(adm)*/

	auth := spotifyStreamer.AuthRPC.GetAuthenticator()

	dealer := cycleBusiness.NewStreamerSuggestionDealer()
	ss, _ := dealer.GetUpdatedMusicList(0, auth)

	fmt.Println(ss)

	cycMan := cycleBusiness.NewCycleManager()
	playGen := playlist.NewPlaylistGenerator(auth)
	
	cycMan.NewListener(playGen)
	cycMan.ManageCycle()



}
