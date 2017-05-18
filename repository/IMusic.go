package repository

import (
	"radiup/cycle"
)

type MusicManager interface {
	RegisterMusic(m Music)
	RemoveMusic(id string)
	SearchMusic(id string)
}