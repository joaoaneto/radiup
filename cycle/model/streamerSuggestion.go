package model

import (
	"time"
)

type StreamerSuggestion struct {
	Musics           []Music
	ModificationDate time.Time
}
