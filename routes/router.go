package routes

import (
	"net/http"

	"groupie-tracker/controllers/handlers"
)

// Routes Handling Automaticly using map
func RoutesHandle(mux *http.ServeMux) {
	// Routes
	Routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":        handlers.HomeHandle,
		"/artists": handlers.ArtistsHandle,
		"/artist":  handlers.ArtistHandle,
		"/aboutUs": handlers.AboutUsHandle,
		"/statics/": handlers.StaticsHandle,
	}
	// for range to handle the routs
	for path, handler := range Routes {
		mux.HandleFunc(path, handler)
	}
}
