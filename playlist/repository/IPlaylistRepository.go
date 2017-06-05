package repository

import (
	"github.com/joaoaneto/radiup/playlist"
)

type PlaylistManager interface {
	Create(p Playlist)
	Update(playlistID int)
	Remove(playlistID int)
	Search(playlistID int) Playlist
}

type PlaylistInfoManager interface {
	Create(p PlaylistInfo)
	Update(userID int, playlistID int)
	Remove(playlistID int)
	Search(playlistID int) PlaylistInfo
}