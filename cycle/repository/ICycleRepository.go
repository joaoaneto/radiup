package repository

import (
	"github.com/joaoaneto/radiup/cycle"
)

type ContentSuggestionManager interface {
	Register(cs cycle.ContentSuggestion)
	Search(nameUser string) []ContentSuggestionRep
}

type CycleManager interface {
	Create(c Cycle)
	Update(id int)
	Remove(id int)
	Search(id int)
}

type MusicManager interface {
	Register(m Music)
	Remove(id string)
	Search(id string)
}

type UserManager interface {
	Create(u User)
	Update(username string)
	Remove(username string)
	Search(username string)
}

type VoluntarySuggestionManager interface {
	Register(v cycle.VoluntarySuggestion)
	Search(nameUser string) []VoluntarySuggestionRep
}