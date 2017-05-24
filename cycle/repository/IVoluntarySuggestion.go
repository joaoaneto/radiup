package repository

import (
	"teste_go/cycle"
)

type IVoluntarySuggestion interface {
	RegisterVSuggestion(v cycle.VoluntarySuggestion)
	//RemoveVSuggestion()
	SearchVSuggestion(nameUser string) []VoluntarySuggestionRep
}
