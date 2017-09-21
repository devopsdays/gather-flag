package service

import (
	"net/http"
)

// Route defines a single route, e.g. a human readable name, HTTP method and the
// pattern the function that will execute when the route is called.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

// Initialize our routes
var routes = Routes{

	Route{
		"GetRoom",         // Name
		"GET",             // HTTP method
		"/rooms/{roomID}", // Route pattern
		GetRoom,
	},
}