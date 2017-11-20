package model

import (
	"time"
)

type VoluntarySuggestion struct {
	Track     Music
	Users     []string
	Timestamp time.Time
	Votes     int
}
