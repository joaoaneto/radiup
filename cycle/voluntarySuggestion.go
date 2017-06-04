package cycle

import (
	"time"
)

type VoluntarySuggestion struct {
	//track Music
	V_User      User
	Votes     int
	Timestamp time.Time
}
