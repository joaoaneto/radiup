package cycle

type ContentSuggestion struct {
	Title                 string
	Description           string
	ContentSuggestionUser User
	Votes                 int
	Validated             bool
	Done                  bool
}
