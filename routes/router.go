package routes

import (
	"net/http"

	"groupietracker/controllers"
)

func RoutesHandle(mux *http.ServeMux) {
	mux.HandleFunc("/", controllers.ArtistsController)
	mux.HandleFunc("/statics/", controllers.ServeStatictsController)
}
