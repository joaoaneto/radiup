package main

import (
	"fmt"
	"github.com/joaoaneto/radiup/streamer"
	"github.com/joaoaneto/radiup/streamer/spotify"
	"net/http"
)

func main() {

	// Improve this, please.
	// I don't found better place for calling this function.
	// By: JN
	spotify.Start()

	spotifyStreamer := streamer.GetStreamerManager().Get("SPOTIFY")

	/*Especified de url of the service*/
	http.HandleFunc("/callback", spotifyStreamer.AuthRPC.NewClientAuth)
	/*http.HandleFunc("/radiup", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Testando")
	})*/

	go http.ListenAndServe(":8080", nil)

	spotifyStreamer.AuthRPC.NewAuthenticator("http://localhost:8080/callback")

	oAuthTest := streamer.OAuthInfo{ClientID: "42d13a4cacae480189b2702e48d6879a", SecretKey: "f0864a30cca443c4b33b940940285d87"}
	spotifyStreamer.AuthRPC.SetAuthInfo(oAuthTest)
	url := spotifyStreamer.AuthRPC.GetAuthURL()

	fmt.Println("Please, use this url for auth:")
	fmt.Println(url)

	// Getting client through channel
	client := <-spotifyStreamer.AuthRPC.GetChannel()

	// Getting client informations
	tkn, _ := client.Token()
	user, _ := client.CurrentUser()

	// This informations should be persist in User AuthInfo
	fmt.Println("User Token:", tkn.AccessToken)
	fmt.Println("User Refresh Token:", tkn.RefreshToken)
	fmt.Println("Token Type:", tkn.TokenType)
	fmt.Println("Expiry Time:", tkn.Expiry)

}
