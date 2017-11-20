package repository

import (
	"time"

	cycle "github.com/joaoaneto/radiup/cycle/model"
	"golang.org/x/oauth2"
)

type StreamerSuggestionManager interface {
	Register(cycleID int, cs cycle.StreamerSuggestion) error
	Update(cycleID int, listMusic []cycle.Music) error
	SearchAll(cycleID int) (cycle.StreamerSuggestion, error)
}

type ContentSuggestionManager interface {
	Register(cycleID int, cs cycle.ContentSuggestion) error
	//Search(nameUser interface{}) ([]cycle.ContentSuggestion, error)
	SearchAll(cycleID int) ([]cycle.ContentSuggestion, error)
}

type CycleManager interface {
	Create(c cycle.Cycle) error
	Update(registeredID int, updatedCycle cycle.Cycle) error
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

type AdminUserManager interface {
	Create(u cycle.AdminUser) error
	Update(user cycle.User,
		authSpotify *oauth2.Token) error
	Remove(username string) error
	Search(username string) (cycle.AdminUser, error)
	SearchAll() ([]cycle.AdminUser, error)
}

type VoluntarySuggestionManager interface {
	Register(cycleID int, vs cycle.VoluntarySuggestion) error
	SearchAll(cycleID int) ([]cycle.VoluntarySuggestion, error)
	Search(cycleID int, musicID string) (cycle.VoluntarySuggestion, error)
}
