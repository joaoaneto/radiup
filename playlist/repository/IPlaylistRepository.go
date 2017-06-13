package repository

import (
	"github.com/joaoaneto/radiup/playlist"
)

type PlaylistManager interface {
	Create(p playlist.Playlist) error
	Update(playlistID int, musics []cycle.Music, cycles cycle.Cycle) error
	Remove(playlistID int) error
	Search(playlistID int) (playlist.Playlist, error)
}