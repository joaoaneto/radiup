package streamer 

import ()

/*RPC interfaces */
type ContentRPC interface {
	GetMusicData()
	GetPlaylistData()
}

type SocialRPC interface {
	GetFollowers()
	GetInstant()
}

type AuthRPC interface {
	NewAuthenticator()
	SetAuthInfo()
	NewClientAuth()
}
