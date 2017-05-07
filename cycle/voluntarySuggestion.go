package cycle

import (
		"fmt"
		"time"
)

type VoluntarySuggestion struct {
	track Music
	user User
	votes int
	timestamp Time
}