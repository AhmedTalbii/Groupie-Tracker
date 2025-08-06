package fetchers

import (
	"groupie-tracker/config"
	"groupie-tracker/models"
)

// fetches artists data from the API using the given ID
// returns a slice of Artist models.
func FetchArtists(id string) []models.Artist {
	var artists []models.Artist
	MustFetch(config.ArtistURL+id, &artists)
	return artists
}

// fetches location data from the API using the given ID
// returns a LOCATIONS model.
func FetchLocaion(id string) models.LOCATIONS {
	var location models.LOCATIONS
	MustFetch(config.LocationURL+id, &location)
	return location
}

// fetches relation data from the API using the given ID
// returns a Relations model.
func FetchRelation(id string) models.Relations {
	var relation models.Relations
	MustFetch(config.RelationsURL+id, &relation)
	return relation
}
