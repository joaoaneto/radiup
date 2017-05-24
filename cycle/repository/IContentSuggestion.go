package repository

import (
	"teste_go/cycle"
)

type IContentSuggestion interface {
	RegisterCSuggestion(cs cycle.ContentSuggestion)
	//RemoveCSuggestion()
	SearchCSuggestion(nameUser string) []ContentSuggestionRep
}
