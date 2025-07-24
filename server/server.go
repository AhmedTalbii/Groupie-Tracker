package server

import (
	"fmt"
	"groupietracker/config"
	"groupietracker/controllers/fetchers"
	"groupietracker/routes"
	"log"
	"net/http"
)

func StartServer() {
	// fetch the data first
	fetchers.InitFetch()

	// declare tha handler
	mux := http.NewServeMux()
	
	// routes handle
	routes.RoutesHandle(mux)

	// server config
	serv := &http.Server{
		Addr: config.Port,
		Handler: mux,
	}

	// print the url then start listening
	fmt.Println("server started at http://localhost"+config.Port)
	log.Fatal(serv.ListenAndServe())
}
