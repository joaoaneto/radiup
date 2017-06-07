package repository

import (
	"github.com/joaoaneto/radiup/playlist"
)

type PlaylistManager interface {
	Create(p playlist.Playlist) string
	Update(playlistID int) string
	Remove(playlistID int) string
	Search(playlistID int) (playlist.Playlist, string)
}

type PlaylistInfoManager interface {
	Create(p playlist.PlaylistInfo) string
	Update(userID int, playlistID int) string
	Remove(playlistID int) string
	Search(playlistID int) (playlist.PlaylistInfo, string)
}
