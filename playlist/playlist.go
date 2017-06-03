package playlist

import (
	"github.com/joaoaneto/radiup/cycle"
)

type Playlist struct {
	PlaylistID int
	Musics []cycle.Music
	Cycle cycle.Cycle
	PlaylistInfo PlaylistInfo
//	PlaylistRPC PlaylistRPC
}