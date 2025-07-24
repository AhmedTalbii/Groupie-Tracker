package fetchers

import (
	"log"

	"groupietracker/config"
	"groupietracker/models"
)

func InitFetch() {
	// fetch the data of artists
	artists, errGettingData := FetchData[[]models.Artist](config.API_ARTISTS_URL)
	if errGettingData != nil {
		log.Fatal("Error Fetching :", errGettingData)
	}
	var pageData []models.ArtistData
	for _, artist := range artists {
		// Locations
		locations := MustFetch[*models.Locations](artist.Locations, "locations")
		// Dates
		dates := MustFetch[*models.Dates](artist.ConcertDates, "dates")
		// Relations
		relations := MustFetch[*models.Relations](artist.Relations, "relations")
		// PrintArtistFullData
		ModalData := &models.ArtistData{
			Id:           artist.Id,
			Image:        artist.Image,
			Name:         artist.Name,
			CreationDate: artist.CreationDate,
			FirstAlbum:   artist.FirstAlbum,
			Members:      artist.Members,
			Concerts:     relations.DatesLocations,
			Locations:    locations.Locations,
			Dates:        dates.Dates,
			Sources:      []string{artist.Locations, artist.ConcertDates, artist.Relations},
		}
		// append to our array where we save all the data 
		pageData = append(pageData, *ModalData)
	}
	// put in the global variable the data
	models.DataFetched = pageData
}
