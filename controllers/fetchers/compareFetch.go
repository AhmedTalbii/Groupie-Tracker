package fetchers

import (
	"strconv"
	"time"

	"groupie-tracker/config"
	"groupie-tracker/helpers"
	"groupie-tracker/models"
)

func CompareAndFetch() {
	if models.LastTimeFetch.Before(time.Now().Add(-config.TimeToRefreshData)) {
		var artists []models.Artist
		var relations models.RelationsIndex
		var locations models.LocationsIndex
		var dates models.DatesIndex

		helpers.Help.Fetch(config.ArtistsURL, &artists)
		helpers.Help.Fetch(config.RelationsURL, &relations)
		helpers.Help.Fetch(config.LocationURL, &locations)
		helpers.Help.Fetch(config.DatesURL, &dates)

		fullData := &models.AllData{
			Artists:   artists,
			Relations: relations,
			Locations: locations,
			Dates:     dates,
		}

		for i, Artist := range fullData.Artists {
			ArtistAllData := &models.ArtistAllData{
				Id:           Artist.Id,
				Image:        Artist.Image,
				Name:         Artist.Name,
				Members:      Artist.Members,
				CreationDate: Artist.CreationDate,
				FirstAlbum:   Artist.FirstAlbum,
				Relations:    fullData.Relations.Index[i].DatesLocations,
				Locations:    fullData.Locations.Index[i].Locations,
				Dates:        fullData.Dates.Index[i].Dates,
				ArtistURL:    "https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(Artist.Id),
				RelationsURL: Artist.RelationsUrl,
				LocationsURL: Artist.LocationsUrl,
				DatesURL:     Artist.DatesUrl,
			}
			models.ArtistsFullData = append(models.ArtistsFullData, *ArtistAllData)
		}

		helpers.Help.RenderPageInsideBuffer(config.Pages+"index", nil, &models.HomeTemplate)
		helpers.Help.RenderPageInsideBuffer(config.Pages+"artists", struct{ Artists []models.ArtistAllData }{Artists: models.ArtistsFullData}, &models.ArtistsTemplate)

		models.LastTimeFetch = time.Now()
	}
}
