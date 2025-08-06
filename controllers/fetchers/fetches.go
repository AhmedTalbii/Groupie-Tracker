package fetchers

import (
	"groupie-tracker/config"
	"groupie-tracker/models"
)

// fetches artists data from the API using the given ID
// returns a slice of Artist models.
func FetchArtists() *[]models.Artist {
	var artists []models.Artist
	MustFetch(config.ArtistURL, &artists)
	return &artists
}

// fetches location data from the API using the given ID
// returns a LOCATIONS model.
func FetchLocaion(id string) *models.Locations {
	var location models.Locations
	MustFetch(config.LocationURL+id, &location)
	return &location
}

// fetches relation data from the API using the given ID
// returns a Relations model.
func FetchRelation(id string) *models.Relations {
	var relation models.Relations
	MustFetch(config.RelationsURL+id, &relation)
	return &relation
}

// fetches Date data from the API using the given ID
// returns a Dates model.
func FetchDates(id string) *models.Dates {
	var dates models.Dates
	MustFetch(config.RelationsURL+id, &dates)
	return &dates
}
