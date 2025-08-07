package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/config"
	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/models"
)

// handles GET requests by fetching all artists and rendering them.
// uses a mutex to safely access shared data between goroutines.
func ArtistsHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// 405 method not allowd
		return
	}

	if len(models.Artists) == 0 {
		models.Mu.Lock()
		models.Artists = *fetchers.FetchArtists()
		models.Templat = template.Must(template.ParseFiles(config.Pages + "artists.html"))
		models.Mu.Unlock()
	}

	models.Templat.Execute(w, struct{ PageData []models.Artist }{PageData: models.Artists})
}
