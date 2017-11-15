package service

import (
	"github.com/gorilla/mux"
	"net/http"
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

	// CRUD

	Route{
		
	},

	Route{
	
	},

	Route{
		
	},

	Route{
	
	},

	//Authentication

	Route{
		
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