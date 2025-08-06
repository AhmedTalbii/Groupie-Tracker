package handlers

import (
	"net/http"

	"groupie-tracker/controllers/rendrers"
)

func AboutUsHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// 405 method not allowd
		return
	}
	// render
	rendrers.MustRender("aboutUs", nil, w)
}
