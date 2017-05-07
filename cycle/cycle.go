package cycle

import (
	"fmt"
	"time"
)

type Cycle struct{
	start Time/*Depois ver como manipula*/
	end Time
	type string
	description string
	voluntarySuggestion VoluntarySuggestion
	//Streamer streamerSuggestions
	contentSuggestion ContentSuggestion
}