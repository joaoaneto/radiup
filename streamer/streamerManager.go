package streamer

type Streamer struct {
	Name    string
	AuthRPC AuthRPC

	/*
		contentRPC ContentRPC
		socialRPC SocialRPC

		oAuthRPC OAuthRPC
	*/
}

type StreamerManager struct {
	Sm map[string]Streamer
}

func (streamerMan *StreamerManager) RegisterStreamer(name string, s Streamer) {
	streamerMan.Sm[name] = s
}

func (streamerMan *StreamerManager) Get(name string) Streamer {
	return streamerMan.Sm[name]
}

/*Variable that will manage de streamers*/

var streamerManager StreamerManager

func GetStreamerManager() StreamerManager {
	return streamerManager
}
