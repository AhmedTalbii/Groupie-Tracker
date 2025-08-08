package handlers

import (
	"net/http"
	"os"

	"groupie-tracker/helpers"
)

// serves static files if the path is valid and not a directory,
// otherwise returns forbidden or server errors.
func StaticsHandle(w http.ResponseWriter, r *http.Request) {
	info, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		helpers.Help.ErrorPage(w, http.StatusNotFound)
		return
	}
	if info.IsDir() {
		helpers.Help.ErrorPage(w, http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
