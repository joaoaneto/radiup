package cycle

import (
)

type ContentSuggestion struct {
	title string
	description string
	user User
	votes int
	validated bool
	done bool
}