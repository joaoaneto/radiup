package model

import (
	"time"
)

type VoluntarySuggestion struct {
	Track                   Music
	VoluntarySuggestionUser User
	Votes                   int
	Timestamp               time.Time
}
