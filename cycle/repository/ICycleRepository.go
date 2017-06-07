package repository

import (
	"time"

	"github.com/joaoaneto/radiup/cycle"
)

type ContentSuggestionManager interface {
	Register(cs cycle.ContentSuggestion) string
	Search(nameUser interface{}) ([]cycle.ContentSuggestion, string)
}

type CycleManager interface {
	Create(c cycle.Cycle)
	Update(registeredID int, start time.Time, end time.Time, cycleType string,
		description string, voluntarySuggestion cycle.VoluntarySuggestion,
		streamerSuggestion cycle.StreamerSuggestion,
		contentSuggestion cycle.ContentSuggestion) string
	Remove(id int) string
	Search(id int) (cycle.Cycle, string)
}

type MusicManager interface {
	Register(m cycle.Music) string
	Remove(id string) string
	Search(id string) (cycle.Music, string)
}

type UserManager interface {
	Create(u cycle.User)
	Update(registered_user string,
		name string,
		password string,
		birth_day time.Time,
		email string,
		sex byte) string
	Remove(username string) string
	Search(username string) (cycle.User, string)
}

type VoluntarySuggestionManager interface {
	Register(v cycle.VoluntarySuggestion) string
	Search(nameUser string) ([]cycle.VoluntarySuggestion, string)
}
