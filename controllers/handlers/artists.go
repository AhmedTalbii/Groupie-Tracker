package handlers

import (
	"net/http"

	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/controllers/rendrers"
	"groupie-tracker/models"
)

// func that fetch and render all the artists
func ArtistsHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// 405 method not allowd
		return
	}
	// fetch
	artists := fetchers.FetchArtists()
	// append all the artists to array of artists
	models.Mu.Lock()
	models.Artists = *artists
	models.Mu.Unlock()
	// render
	rendrers.MustRender("artists", struct{ PageData []models.Artist }{PageData: *artists}, w)
}
