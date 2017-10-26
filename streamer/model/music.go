package model

import "time"

type Music struct {
	Name     string
	Artist   []string
	ID       string
	SourceID int
	PlayedAt time.Time
}
