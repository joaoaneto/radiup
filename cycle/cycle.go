package cycle

import (
	"fmt"
	"time"
)

type Cycle struct{
	start Time
	end Time
	type string
	description string
	voluntarySuggestion VoluntarySuggestion
	StreamerSuggestion StreamerSuggestion
	contentSuggestion ContentSuggestion
}