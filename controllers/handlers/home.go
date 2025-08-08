package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/config"
	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/helpers"
	"groupie-tracker/models"
)

// serves the homepage by validating the request path, connection and method,
// then rendering the "index" template if it's a valid GET request to "/".
// fetch artists data before moving to the next page
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		helpers.Help.ErrorPage(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		helpers.Help.ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}
	var err error

	_, err = http.Get("https://www.google.com")
	if err != nil {
		helpers.Help.ErrorPage(w, http.StatusInternalServerError)
		return
	}

	helpers.Help.RenderPage(config.Pages+"index", nil, w)

	if len(models.Artists) == 0 {
		models.Mu.Lock()
		models.Artists = *fetchers.FetchArtists()
		models.Templat, err = template.ParseFiles(config.Pages + "artists.html")
		models.Mu.Unlock()
		if err != nil {
			helpers.Help.ErrorPage(w, http.StatusInternalServerError)
			return
		}
	}
}
