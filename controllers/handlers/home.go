package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/config"
	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/controllers/rendrers"
	"groupie-tracker/models"
)

// serves the homepage by validating the request path, connection and method,
// then rendering the "index" template if it's a valid GET request to "/".
// fetch artists data before moving to the next page
func HomeHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		rendrers.ErrorPage(models.Data{Error: "Page Not Found", StatusE: "404"}, w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		rendrers.ErrorPage(models.Data{Error: "Method Not Allowed", StatusE: "405"}, w, http.StatusMethodNotAllowed)
		return
	}
	var err error

	_, err = http.Get("https://www.google.com")
	if err != nil {
		rendrers.ErrorPage(models.Data{Error: "Connection Error", StatusE: "500"}, w, http.StatusInternalServerError)
		return
	}
	rendrers.MustRender("index", nil, w)

	if len(models.Artists) == 0 {
		models.Mu.Lock()
		defer models.Mu.Unlock()

		models.Artists = *fetchers.FetchArtists(w)
		if len(models.Artists) == 0 {
			rendrers.ErrorPage(models.Data{Error: "Error Internal Server", StatusE: "500"}, w, http.StatusInternalServerError)
			return
		}
		models.Templat, err = template.ParseFiles(config.Pages + "artists.html")
		if err != nil {
			rendrers.ErrorPage(models.Data{Error: "Error Internal Server", StatusE: "500"}, w, http.StatusInternalServerError)
			return
		}
	}
}
