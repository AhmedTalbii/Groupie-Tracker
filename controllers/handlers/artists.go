package handlers

import (
	"html/template"
	"net/http"

	"groupie-tracker/config"
	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/controllers/rendrers"
	"groupie-tracker/models"
)

// handles GET requests by fetching all artists and rendering them.
// uses a mutex to safely access shared data between goroutines.
// in case the is a fetching problem will render the error page
func ArtistsHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rendrers.ErrorPage(models.Data{Error: "Method Not Allowed", StatusE: "405"}, w, http.StatusMethodNotAllowed)
		return
	}
	var err error
	if len(models.Artists) == 0 {
		
		models.Mu.Lock()
		defer models.Mu.Unlock()

		models.Artists = *fetchers.FetchArtists()
		models.Templat, err = template.ParseFiles(config.Pages + "artists.html")
		if err != nil{
			rendrers.ErrorPage(models.Data{Error: "Error Internal Server", StatusE: "500"}, w, http.StatusInternalServerError)
			return
		}
	}

	if len(models.Artists) == 0{
		rendrers.ErrorPage(models.Data{Error: "Error Internal Server", StatusE: "500"}, w, http.StatusInternalServerError)
		return
	}
	models.Templat.Execute(w, struct{ PageData []models.Artist }{PageData: models.Artists})
}
