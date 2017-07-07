package cycle

import (
	"time"
)

// Music is the song type on RadiUp!
type Music struct {
	Name     string
	Artist   []string
	ID       string
	SourceID int
	PlayedAt time.Time
}
