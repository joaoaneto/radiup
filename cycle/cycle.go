package cycle

import (
	"time"
)

type Cycle struct{
	id int	// temp
	start time.Time
	end time.Time
	//type string
	description string
	voluntarySuggestion VoluntarySuggestion
	StreamerSuggestion streamerSuggestion
	contentSuggestion ContentSuggestion
}