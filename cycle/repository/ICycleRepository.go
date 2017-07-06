package repository

import (
	"time"

	"github.com/joaoaneto/radiup/cycle"
)

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

type UserManager interface {
	Create(u cycle.User) error
	Update(registered_user string,
		name string,
		password string,
		birth_day time.Time,
		email string,
		sex byte) error
	Remove(username string) error
	Search(username string) (cycle.User, error)
}

type VoluntarySuggestionManager interface {
	Register(v cycle.VoluntarySuggestion) error
	Search(nameUser string) ([]cycle.VoluntarySuggestion, error)
}
