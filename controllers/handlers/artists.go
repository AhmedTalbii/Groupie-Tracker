package handlers

import (
	"net/http"

	"groupie-tracker/biblio"
	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/models"
)

// Handles GET requests to fetch all artists, render them, and show an error page if fetching fails. Uses a mutex for safe concurrent data access.
func ArtistsHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		biblio.Help.ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}

	if biblio.Help.CheckConnection() != nil{
		biblio.Help.ErrorPage(w,http.StatusServiceUnavailable)
		return
	}
	
	models.Mu.Lock()
	if err := fetchers.InitFetch(); err != nil {
		biblio.Help.ErrorPage(w, http.StatusInternalServerError)
		return
	}
	models.Mu.Unlock()
	w.Write(models.ArtistsTemplate.Bytes())
}
