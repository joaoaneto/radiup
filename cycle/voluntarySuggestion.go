package cycle

import (
	"time"
)

type VoluntarySuggestion struct {
	//track Music
	user      User
	votes     int
	Timestamp time.Time
}
