package server

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/config"
	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/routes"
)

// initializes the HTTP server by setting up routes,
// configuring the server with the specified port and handler,
// then starts listening and serving requests.
func StartServer() {
	fetchers.InitFetch()

	mux := http.NewServeMux()
	routes.RoutesHandle(mux)
	serv := &http.Server{
		Addr:    config.Port,
		Handler: mux,
	}
	fmt.Println("server started at http://localhost" + config.Port)
	log.Fatal(serv.ListenAndServe())
}
