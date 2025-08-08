package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/config"
	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/controllers/rendrers"
	"groupie-tracker/models"
)

// processes GET requests by extracting the artist ID from the URL,
// validating it, fetching related data (locations, dates, relations),
// and rendering the artist detail page using the combined data.
func ArtistHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rendrers.ErrorPage(models.Data{Error: "Method Not Allowed", StatusE: "405"}, w, http.StatusMethodNotAllowed)
		return
	}
	Url := strings.Split(r.URL.Path, "/")
	Id := Url[len(Url)-1]
	idInt, err := strconv.Atoi(Id)

	if len(models.Artists) == 0 {
		models.Mu.Lock()
		models.Artists = *fetchers.FetchArtists()
		models.Templat = template.Must(template.ParseFiles(config.Pages + "artists.html"))
		models.Mu.Unlock()
	}

	if err != nil || idInt > len(models.Artists) || idInt <= 0 {
		rendrers.ErrorPage(models.Data{Error: "Page Not Found", StatusE: "404"}, w, http.StatusNotFound)
		return
	}

	fullData := &models.FullArtistsData{
		Artist:    &models.Artists[idInt-1],
		Locations: fetchers.FetchLocaion(Id),
		Dates:     fetchers.FetchDates(Id),
		Relations: fetchers.FetchRelation(Id),
	}
	rendrers.MustRender("artist", fullData, w)
}
