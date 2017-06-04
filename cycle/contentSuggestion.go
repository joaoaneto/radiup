package cycle

import (
)


type ContentSuggestion struct {
	Title string
	Description string
	_User User
	Votes int
	Validated bool
	Done bool
}