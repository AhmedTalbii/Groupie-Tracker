package controllers

import (
	"net/http"

	"groupietracker/controllers/renderers"
	"groupietracker/models"
)

func ArtistsController(w http.ResponseWriter, r *http.Request) {
	pageDataRender := models.PageData{
		Artists: models.DataFetched,
	}
	renderers.PageRender(w, r, "index.html", pageDataRender)
}
