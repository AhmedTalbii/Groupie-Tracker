package server

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/config"
	"groupie-tracker/routes"
)

// This func is the main function that starts our server
func StartServer() {
	// declare tha handler
	mux := http.NewServeMux()

	// routes handle
	routes.RoutesHandle(mux)

	// server config
	serv := &http.Server{
		Addr:    config.Port,
		Handler: mux,
	}

	// print the url then start listening
	fmt.Println("server started at http://localhost" + config.Port)
	log.Fatal(serv.ListenAndServe())
}
