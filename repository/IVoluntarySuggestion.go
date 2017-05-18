package repository

import (
	"radiup/cycle"
)

//Does it need updates?
type VSuggestionManager interface {
	RegisterVSuggestion()
	RemoveVSuggestion()
	SearchVSuggestion()
}