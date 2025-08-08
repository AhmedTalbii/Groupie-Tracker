package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/config"
	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/helpers"
	"groupie-tracker/models"
)

// processes GET requests by extracting the artist ID from the URL,
// validating it, fetching related data (locations, dates, relations),
// and rendering the artist detail page using the combined data.
func ArtistHandle(w http.ResponseWriter, r *http.Request) {
	helpers.Help.CheckGet(w, r)

	Url := strings.Split(r.URL.Path, "/")
	Id := Url[len(Url)-1]
	idInt, err := strconv.Atoi(Id)

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

	if err != nil || idInt > len(models.Artists) || idInt <= 0 {
		helpers.Help.ErrorPage(w, http.StatusNotFound)
		return
	}

	fullData := &models.FullArtistsData{
		Artist:    &models.Artists[idInt-1],
		Locations: fetchers.FetchLocaion(Id),
		Dates:     fetchers.FetchDates(Id),
		Relations: fetchers.FetchRelation(Id),
	}
	
	helpers.Help.RenderPage(config.Pages+"artist", fullData, w)
}
