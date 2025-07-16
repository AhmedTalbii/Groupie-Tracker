package server

import (
	"fmt"
	"groupietracker/config"
	"groupietracker/routes"
	"log"
	"net/http"
)

func StartServer() {
	mux := http.NewServeMux()
	
	// routes handle
	routes.RoutesHandle(mux)

	serv := &http.Server{
		Addr: config.Port,
		Handler: mux,
	}

	fmt.Println("server started at http://localhost"+config.Port)
	log.Fatal(serv.ListenAndServe())
}
