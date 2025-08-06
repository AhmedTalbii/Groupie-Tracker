package handlers

import (
	"net/http"

	"groupie-tracker/controllers/rendrers"
)

// Handle the main page
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// 404 Page Not Found
		return
	}
	if r.Method != http.MethodGet {
		// 405 method not allowd
		return
	}
	// render

	rendrers.MustRender("index", nil, w)
}
