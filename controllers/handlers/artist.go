package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/controllers/fetchers"
	"groupie-tracker/controllers/rendrers"
	"groupie-tracker/models"
)

// func that fetch and render all the artist
func ArtistHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// 405 method not allowd
		return
	}
	// id from url
	Url := strings.Split(r.URL.Path, "/")
	Id := Url[len(Url)-1]
	idInt, err := strconv.Atoi(Id)
	if err != nil || idInt > len(models.Artists) || idInt <= 0 {
		// 400 Not Found
		return
	}

	// define full data
	fullData := &models.FullArtistsData{
		Artist:    &models.Artists[idInt-1],
		Locations: fetchers.FetchLocaion(Id),
		Dates:     fetchers.FetchDates(Id),
		Relations: fetchers.FetchRelation(Id),
	}

	// render
	rendrers.MustRender("arist", fullData, w)
}
