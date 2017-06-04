package cycle

import (
)


type ContentSuggestion struct {
	Title 				  string
	Description 		  string
	ContentSuggestionUser User
	Votes 				  int
	Validated 			  bool
	Done 				  bool

}