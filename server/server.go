package server

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/config"
	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/routes"
)

// start fetching Data 
// Starts HTTP server with the given routes, port, and handler, then begins serving requests.
func StartServer() {
	if err := fetchers.InitFetch(); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	routes.RoutesHandle(mux)
	serv := &http.Server{
		Addr:    config.Port,
		Handler: mux,
	}
	fmt.Println("server started at http://localhost" + config.Port)
	log.Fatal(serv.ListenAndServe())
}
