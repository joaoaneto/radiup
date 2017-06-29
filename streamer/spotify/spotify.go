package spotify

import (
	"github.com/joaoaneto/radiup/streamer"
)

const kStreamerSpotify string = "SPOTIFY"

/*Quando implementar as classes do Wrapper que irão definir as interfaces, atribuir ela aqui*/
func newStreamerSpotify() streamer.Streamer {
	authSpotify := NewAuthSpotify()
	socialSpotify := NewSocialSpotify()
	spotifyStreamer := streamer.Streamer{Name: kStreamerSpotify, AuthRPC: authSpotify, SocialRPC: socialSpotify}
	return spotifyStreamer
}

func Start() {
	streamer.GetStreamerManager().RegisterStreamer(kStreamerSpotify, newStreamerSpotify())
}
