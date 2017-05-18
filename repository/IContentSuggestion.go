package repository

import (
	"radiup/cycle"
)

//Does it need updates?
type CSuggestionManager interface {
	RegisterCSuggestion()
	RemoveCSuggestion()
	SearchCSuggestion()
}