package business

import (
	"github.com/joaoaneto/radiup/cycle"
	"github.com/zmb3/spotify"
)

type StreamerSuggestionOperator interface {
	GetUpdatedMusicList(auth spotify.Authenticator) (cycle.StreamerSuggestion, error)
}
