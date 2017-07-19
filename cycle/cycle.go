package cycle

import (
	"time"
)

type Cycle struct {
	ID                       int // temp
	Start                    time.Time
	End                      time.Time
	CycleType                string
	Description              string
	CycleVoluntarySuggestion []VoluntarySuggestion
	CycleStreamerSuggestion  StreamerSuggestion
	CycleContentSuggestion   []ContentSuggestion
}