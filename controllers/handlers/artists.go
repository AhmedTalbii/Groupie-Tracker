package handlers

import (
	"net/http"

	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/helpers"
	"groupie-tracker/models"
)

// handles GET requests by fetching all artists and rendering them.
// uses a mutex to safely access shared data between goroutines.
// in case the is a fetching problem will render the error page
func ArtistsHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		helpers.Help.ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}

	models.Mu.Lock()
	fetchers.CompareAndFetch()
	models.Mu.Unlock()

	w.Write(models.ArtistsTemplate.Bytes())
}
