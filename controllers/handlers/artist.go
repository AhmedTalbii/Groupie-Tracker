package handlers

import (
	"net/http"
	"strconv"
	"strings"

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
	if err != nil || idInt > len(models.Artists) || idInt <= 0 {
		rendrers.ErrorPage(models.Data{Error: "Page Not Found", StatusE: "404"}, w, http.StatusNotFound)
		return
	}

	fullData := &models.FullArtistsData{
		Artist:    &models.Artists[idInt-1],
		Locations: fetchers.FetchLocaion(Id, w),
		Dates:     fetchers.FetchDates(Id, w),
		Relations: fetchers.FetchRelation(Id, w),
	}
	rendrers.MustRender("artist", fullData, w)
}
