package spotify

import ("radiup/streamer")

const kStreamerSpotify := "SPOTIFY"

/*Quando implementar as classes do Wrapper que irão definir as interfaces, atribuir ela aqui*/
func newStreamerSpotify() streamer{
	spotifyStreamer := streamer{name:kStreamerSpotify}
	return spotifyStreamer
}
/*Registra streamer spotify*/
func init(){
	GetStreamerManager().RegisterStreamer(kStreamerSpotify, newStreamerSpotify)
}	