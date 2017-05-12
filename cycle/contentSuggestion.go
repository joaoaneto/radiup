package cycle

import (
	"fmt"
)

type ContentSuggestion struct {
	title string
	description string
	user User
	votes int
	validated bool
	done bool
}