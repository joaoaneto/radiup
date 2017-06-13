package playlist

import (
	"github.com/joaoaneto/radiup/cycle"
)

type Playlist struct {
	PlaylistID int
	Musics []cycle.Music
	Cycles cycle.Cycle
//	PlaylistRPC PlaylistRPC		??????
}