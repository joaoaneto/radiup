package streamer 

import()

type Streamer struct{
	name string
	/*
	contentRPC ContentRPC
	socialRPC SocialRPC
	authRPC AuthRPC
	oAuthRPC OAuthRPC
	*/
} 

type StreamerManager struct{
	Sm map[string]Streamer
}

func (streamerMan *StreamerManager) RegisterStreamer(name string, s Streamer){
	streamerMan.Sm[name] = s
}

func (streamerMan *StreamerManager) Get(name string) Streamer{
	return streamerMan.Sm[name]
}

/*Variable that will manage de streamers*/

var streamerManager StreamerManager

func GetStreamerManager() StreamerManager{
	return streamerManager
}