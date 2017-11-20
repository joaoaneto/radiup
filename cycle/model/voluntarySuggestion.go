package model

import (
	"time"
)

type VoluntarySuggestion struct {
	Track                    Music
	VoluntarySuggestionUsers []User
	Timestamp                time.Time
	/*Votes                   int*/
}
