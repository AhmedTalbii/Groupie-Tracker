package handlers

import (
	"net/http"

	"groupie-tracker/controllers/rendrers"
)

// serves the homepage by validating the request path and method,
// then rendering the "index" template if it's a valid GET request to "/".
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// 404 Page Not Found
		return
	}
	if r.Method != http.MethodGet {
		// 405 method not allowd
		return
	}
	rendrers.MustRender("index", nil, w)
}
