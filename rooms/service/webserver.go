package service

import (
	"log"
	"net/http"
)

// StartWebServer starts the HTTP server for the service
func StartWebServer(port string) {

	r := NewRouter()
	http.Handle("/", r)

	log.Println("Starting HTTP service at " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
