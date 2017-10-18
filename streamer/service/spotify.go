package service

import "github.com/joaoaneto/radiup/streamer/service/spotify"

const kStreamerSpotify string = "SPOTIFY"

/*Quando implementar as classes do Wrapper que irão definir as interfaces, atribuir ela aqui*/
func newStreamerSpotify() Streamer {
	authSpotify := spotify.NewAuthSpotify()
	socialSpotify := spotify.NewSocialSpotify()
	contentSpotify := spotify.NewContentSpotify()
	spotifyStreamer := Streamer{Name: kStreamerSpotify, AuthRPC: authSpotify, SocialRPC: socialSpotify, ContentRPC: contentSpotify}
	return spotifyStreamer
}

func Start() {
	GetStreamerManager().RegisterStreamer(kStreamerSpotify, newStreamerSpotify())
}
