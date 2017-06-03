package playlist

import (
	"github.com/joaoaneto/radiup/cycle"
)

type PlaylistRPC interface {
	AddTrack(track cycle.Music)
	DeleteTrack(track cycle.Music)
}