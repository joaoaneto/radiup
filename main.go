package main

import (
	//"fmt"
	"net/http"

	"github.com/joaoaneto/radiup/cycle/controller"
	"github.com/joaoaneto/radiup/streamer"
	"github.com/joaoaneto/radiup/streamer/spotify"
)

func main() {

	// Improve this, please.
	// I don't found better place for calling this function.
	// By: JN
	spotify.Start()

	spotifyStreamer := streamer.GetStreamerManager().Get("SPOTIFY")

	/*spotifyStreamer.AuthRPC.NewAuthenticator("http://localhost:8080/register")
	oAuthTest := streamer.OAuthInfo{ClientID: "42d13a4cacae480189b2702e48d6879a", SecretKey: "f0864a30cca443c4b33b940940285d87"}
	spotifyStreamer.AuthRPC.SetAuthInfo(oAuthTest)*/

	http.HandleFunc("/callback", spotifyStreamer.AuthRPC.NewClientAuth)
	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/register", controller.RegisterHandler)
	http.HandleFunc("/content/list", controller.ShowContentSuggestionsHandler)
	http.HandleFunc("/content/register", controller.RegisterContentSuggestionsHandler)
	http.HandleFunc("/logout", controller.LogoutHandler)

	http.ListenAndServe(":8080", nil)

	//RunServer()

	//url := spotifyStreamer.AuthRPC.GetAuthURL()

	//fmt.Println("Please, use this url for auth:")
	//fmt.Println(url)

	// Getting client through channel
	/*client := <-spotifyStreamer.AuthRPC.GetChannel()

	// Getting client informations
	tkn, _ := client.Token()

	user, _ := client.CurrentUser()

	// This informations should be persist in User AuthInfo
	fmt.Println("User Token:", tkn.AccessToken)
	fmt.Println("User Refresh Token:", tkn.RefreshToken)
	fmt.Println("Token Type:", tkn.TokenType)
	fmt.Println("Expiry Time:", tkn.Expiry)

	// Info about User
	fmt.Println("User Name:", user.DisplayName)
	fmt.Println("User ID:", user.ID)
	fmt.Println("User Birthdate:", user.Birthdate)

	music, musicErr := spotifyStreamer.SocialRPC.GetInstant(client)

	// Info about the music the user is currently listening to
	if musicErr == nil {
		fmt.Println("Music Name: ", music.Name)
		fmt.Println("Music Artists: ", music.Artist)
		fmt.Println("Music ID: ", music.ID)
	}

	listMusic, _ := spotifyStreamer.SocialRPC.GetLastPlayedMusics(client)

	// Info about the recently played songs
	for _, music := range listMusic {
		fmt.Println("Music Name: ", music.Name)
		fmt.Println("Music Artists: ", music.Artist)
		fmt.Println("Music ID: ", music.ID)
		fmt.Println("Music Played at: ", music.PlayedAt)
	}

	//Info about playlist
	//playlist, _ := spotifyStreamer.ContentRPC.GetPlaylistData(client)
	//fmt.Println(playlist)

	//Search test
	search, _ := spotifyStreamer.ContentRPC.GetMusicData(client, "Show das poderosas")

	for _, i := range search.Tracks.Tracks {
		fmt.Println(i.SimpleTrack)
	} */

}
