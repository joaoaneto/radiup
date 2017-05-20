package cycle

import (
	"fmt"
	"time"
)

type Cycle struct{
	id int	// temp
	start Time
	end Time
	type string
	description string
	voluntarySuggestion VoluntarySuggestion
	StreamerSuggestion StreamerSuggestion
	contentSuggestion ContentSuggestion
}