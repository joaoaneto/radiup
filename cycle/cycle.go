package cycle

import (
	"time"
)

type Cycle struct{
	Id int	// temp
	Start time.Time
	End time.Time
	//type string
	Description string
	VoluntarySuggestion VoluntarySuggestion
	StreamerSuggestion streamerSuggestion
	ContentSuggestion ContentSuggestion
}