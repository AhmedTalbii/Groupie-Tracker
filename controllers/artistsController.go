package controllers

import (
	"log"
	"net/http"

	"groupietracker/config"
	"groupietracker/controllers/fetchers"
	"groupietracker/controllers/renderers"
	"groupietracker/models"
)

func ArtistsController(w http.ResponseWriter, r *http.Request) {
	// fetch the data of artists
	Artists, errGettingData := fetchers.FetchData(config.API_ARTISTS_URL)
	if errGettingData != nil {
		log.Fatal("Error Fetching :", errGettingData)
	}

	pageData := models.PageDataArtists{
		Artists: Artists,
	}
	
	renderers.PageRender(w, r, "index.html", pageData)
}
