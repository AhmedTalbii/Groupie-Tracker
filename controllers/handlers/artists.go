package handlers

import (
	"net/http"

	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/controllers/rendrers"
	"groupie-tracker/models"
)

// handles GET requests by fetching all artists and rendering them.
// uses a mutex to safely access shared data between goroutines.
func ArtistsHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// 405 method not allowd
		return
	}
	models.Mu.Lock()
	models.Artists = *fetchers.FetchArtists()
	models.Mu.Unlock()
	rendrers.MustRender("artists", struct{ PageData []models.Artist }{PageData: models.Artists}, w)
}
