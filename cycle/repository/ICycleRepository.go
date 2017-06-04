package repository

import (
	"github.com/joaoaneto/radiup/cycle"
)

type ContentSuggestionManager interface {
	Register(cs cycle.ContentSuggestion)
	Search(nameUser interface{}) []ContentSuggestion
}

type CycleManager interface {
	Create(c Cycle)
	Update(registered_id int, start time.Time, end time.Time, _type string,
				   description string, voluntarySuggestion cycle.VoluntarySuggestion,
				   streamerSuggestion cycle.StreamerSuggestion,
				   contentSuggestion cycle.ContentSuggestion)
	Remove(id int)
	Search(id int) cycle.Cycle
}

type MusicManager interface {
	Register(m cycle.Music)
	Remove(id string)
	Search(id string) cycle.Music
}

type UserManager interface {
	Create(u cycle.User)
	Update(registered_user string,
								  name string,
								  password string,
								  birth_day time.Time,
								  email string,
								  sex byte)
	Remove(username string)
	Search(username string) cycle.User
}

type VoluntarySuggestionManager interface {
	Register(v cycle.VoluntarySuggestion)
	Search(nameUser string) []VoluntarySuggestion
}