package handlers

import (
	"net/http"

	"groupie-tracker/biblio"
)

// Serves static files if valid; returns error pages for invalid paths or directories.
func StaticsHandle(w http.ResponseWriter, r *http.Request) {
	if biblio.Help.CheckConnection() != nil {
		biblio.Help.ErrorPage(w, http.StatusServiceUnavailable)
		return
	}
	biblio.Help.StaticsHandler(w, r)
}
