package streamer

import ()

const kStreamerSpotify := "SPOTIFY"

/*This is not ok, i think so*/
func newStreamerSpotify() streamer{
	spotifyStreamer := streamer{name:kStreamerSpotify}
	return spotifyStreamer
}
/*Registra s*/
func init(){
	GetStreamerManager().RegisterStreamer(kStreamerSpotify, newStreamerSpotify)
}	