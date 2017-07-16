package repository

import (
	"time"

	"github.com/joaoaneto/radiup/cycle"
	"golang.org/x/oauth2"
)

type StreamerSuggestionManager interface {
	Register(cs cycle.StreamerSuggestion) error
	SearchAll() ([]cycle.StreamerSuggestion, error)
}

type ContentSuggestionManager interface {
	Register(cs cycle.ContentSuggestion) error
	Search(nameUser interface{}) ([]cycle.ContentSuggestion, error)
	SearchAll() ([]cycle.ContentSuggestion, error)
}

type CycleManager interface {
	Create(c cycle.Cycle) error
	Update(registeredID int, start time.Time, end time.Time, cycleType string,
		description string, voluntarySuggestion cycle.VoluntarySuggestion,
		streamerSuggestion cycle.StreamerSuggestion,
		contentSuggestion cycle.ContentSuggestion) error
	Remove(id int) error
	Search(id int) (cycle.Cycle, error)
}

type MusicManager interface {
	Register(m cycle.Music) error
	Remove(id string) error
	Search(id string) (cycle.Music, error)
}

type SimpleUserManager interface {
	Create(u cycle.SimpleUser) error
	Update(registered_user string,
		name string,
		password []byte,
		birth_day time.Time,
		email string,
		sex string,
		authSpotify *oauth2.Token) error
	Remove(username string) error
	Search(username string) (cycle.SimpleUser, error)
	SearchAll() ([]cycle.SimpleUser, error)
}

type VoluntarySuggestionManager interface {
	Register(v cycle.VoluntarySuggestion) error
	Search(nameUser string) ([]cycle.VoluntarySuggestion, error)
}
