package main 

import (
	mux "github.com/julienschmidt/httprouter"
)

// Struct of route data
type Route struct {
	Method 	string
	Path	string
	Handle 	mux.Handle
}

// Slice to hold Route structs
type Routes []Route

var routes = Routes{

	Route{
		"GET",
		"/",
		Index,
	},

	Route{
		"GET",
		"/houses",
		HouseIndex,
	},

	Route{
		"POST",
		"/houses",
		HouseCreate,
	},

	Route{
		"GET",
		"/houses/id/:id",
		HouseById,
	},

	Route{
		"GET",
		"/houses/city/:city",
		HouseByCity,
	},

	Route{
		"GET",
		"/houses/state/:state",
		HouseByState,
	},
}