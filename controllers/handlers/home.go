package handlers

import (
	"net/http"

	"groupie-tracker/biblio"
	"groupie-tracker/models"
)

// Serves the homepage, validating path/method and rendering the "index" template.
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if biblio.Help.CheckConnection() != nil{
		biblio.Help.ErrorPage(w,http.StatusServiceUnavailable)
		return
	}
	if r.URL.Path != "/" {
		biblio.Help.ErrorPage(w, http.StatusNotFound)
		return
	}
	
	if r.Method != http.MethodGet {
		biblio.Help.ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write(models.HomeTemplate.Bytes())
}
