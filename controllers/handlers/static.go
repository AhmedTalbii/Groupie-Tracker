package handlers

import (
	"net/http"
	"os"

	"groupie-tracker/controllers/rendrers"
	"groupie-tracker/models"
)

// serves static files if the path is valid and not a directory,
// otherwise returns forbidden or server errors.
func StaticsHandle(w http.ResponseWriter, r *http.Request) {
	info, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		rendrers.ErrorPage(models.Data{Error: "Page Not Found",StatusE: "404"},w,http.StatusNotFound)
		return
	}
	if info.IsDir() {
		rendrers.ErrorPage(models.Data{Error: "Error Forbidden", StatusE: "403"}, w, http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, r.URL.Path[1:])
}
