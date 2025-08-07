package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/config"
	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/controllers/rendrers"
	"groupie-tracker/models"
)

// serves the homepage by validating the request path and method,
// then rendering the "index" template if it's a valid GET request to "/".
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// 404 Page Not Found
		return
	}
	if r.Method != http.MethodGet {
		// 405 method not allowd
		return
	}
	rendrers.MustRender("index", nil, w)

	if len(models.Artists) == 0 {
		models.Mu.Lock()
		models.Artists = *fetchers.FetchArtists()
		models.Templat = template.Must(template.ParseFiles(config.Pages + "artists.html"))
		models.Mu.Unlock()
	}
}
