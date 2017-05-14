package streamer 

import()

type streamer struct{
	name string
	contentRPC ContentRPC
	socialRPC SocialRPC
	authRPC AuthRPC
	oAuthRPC OAuthRPC
} 

type StreamerManager struct{
	sm map[string]streamer
}

func (streamerMan *StreamerManager) RegisterStreamer(name string, s streamer){
	streamerMan.sm[name] = streamer
}

func (streamerMan *StreamerManager) Get(name string) streamer{
	return streamerMan.sm[name]
}

/*Variable that will manage de streamers*/
streamerManager StreamerManager

func GetStreamerManager() StreamerManager{
	return streamerManager
}