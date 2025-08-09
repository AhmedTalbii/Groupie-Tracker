package routes

import (
	"net/http"

	"groupie-tracker/controllers/handlers"
)

// Automatically registers routes from a map to the provided ServeMux.
func RoutesHandle(mux *http.ServeMux) {
	Routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":         handlers.HomeHandle,
		"/artists":  handlers.ArtistsHandle,
		"/statics/": handlers.StaticsHandle,
	}
	for path, handler := range Routes {
		mux.HandleFunc(path, handler)
	}
}
