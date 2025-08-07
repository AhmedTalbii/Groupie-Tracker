package fetchers

import (
	"net/http"

	"groupie-tracker/config"
	"groupie-tracker/models"
)

// fetches artists data from the API using the given ID
// returns a slice of Artist models.
func FetchArtists(w http.ResponseWriter) *[]models.Artist {
	var artists []models.Artist
	MustFetch(config.ArtistURL, &artists, w)
	return &artists
}

// fetches location data from the API using the given ID
// returns a LOCATIONS model.
func FetchLocaion(id string, w http.ResponseWriter) *models.Locations {
	var location models.Locations
	MustFetch(config.LocationURL+id, &location, w)
	return &location
}

// fetches relation data from the API using the given ID
// returns a Relations model.
func FetchRelation(id string, w http.ResponseWriter) *models.Relations {
	var relation models.Relations
	MustFetch(config.RelationsURL+id, &relation, w)
	return &relation
}

// fetches Date data from the API using the given ID
// returns a Dates model.
func FetchDates(id string, w http.ResponseWriter) *models.Dates {
	var dates models.Dates
	MustFetch(config.RelationsURL+id, &dates,w)
	return &dates
}
