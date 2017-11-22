package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

// Set of Routes/Endpoints that make an API REST

var routes = Routes{

	Route{
		"RegisterVoluntarySuggestion",
		"POST",
		"/suggestion",
		RegisterVoluntarySuggestionHandler,
	},

	Route{
		"GetVoluntarySuggestion",
		"GET",
		"/suggestion/list",
		GetVoluntarySuggestionHandler,
	},
	// GET localhost:6969/suggestion/vote/username=netoax&mid=123i12ijasdija
	// Search: how to get parameters go http
	Route{
		"RegisterVote",
		"GET",
		"/suggestion/vote",
		RegisterVoteHandler,
	},
}

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {

		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)

	}

	return router

}
