package controllers

import (
	"net/http"

	"groupietracker/controllers/fetchers"
	"groupietracker/controllers/renderers"
	"groupietracker/models"
)

func ArtistsController(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderers.ErrorRendrer(w, r, http.StatusNotFound, " Page Not Found")
		return
	}

	// Time to live check for Update 
	fetchers.TimeToLiveFetch()

	pageDataRender := models.PageData{
		Artists: models.DataFetched,
	}
	renderers.PageRender(w, r, "index.html", pageDataRender)
}
