package handlers

import (
	"net/http"

	"groupie-tracker/controllers/rendrers"
)

// handles GET requests and renders the "AboutUs" page.
func AboutUsHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// 405 method not allowd
		return
	}
	rendrers.MustRender("aboutUs", nil, w)
}
