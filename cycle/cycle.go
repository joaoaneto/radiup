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
	CycleVoluntarySuggestion VoluntarySuggestion
	CycleStreamerSuggestion StreamerSuggestion
	CycleContentSuggestion ContentSuggestion
}