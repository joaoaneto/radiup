package cycle

import (
	"time"
)

type VoluntarySuggestion struct {
	//track Music
	VoluntarySuggestionUser      User
	Votes     					 int
	Timestamp 					 time.Time

}
