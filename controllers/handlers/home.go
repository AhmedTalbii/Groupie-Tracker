package handlers

import (
	"net/http"

	"groupie-tracker/helpers"
	"groupie-tracker/models"
)

// serves the homepage by validating the request path, connection and method,
// then rendering the "index" template if it's a valid GET request to "/".
// fetch artists data before moving to the next page
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		helpers.Help.ErrorPage(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		helpers.Help.ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}

	w.Write(models.HomeTemplate.Bytes())
}
