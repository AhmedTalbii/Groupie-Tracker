package routes

import (
	"net/http"

	"groupietracker/controllers"
)

func RoutesHandle(mux *http.ServeMux) {
	mux.HandleFunc("/", controllers.ArtistsController)
	mux.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))))
}
