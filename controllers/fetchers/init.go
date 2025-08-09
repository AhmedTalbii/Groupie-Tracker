package fetchers

import (
	"strconv"
	"sync"
	"time"

	"groupie-tracker/biblio"
	"groupie-tracker/config"
	"groupie-tracker/models"
)

func InitFetch() error {
	if models.LastTimeFetch.Before(time.Now().Add(-config.TimeToRefreshData)) {
		var artists []models.Artist
		var relations models.RelationsIndex
		var locations models.LocationsIndex
		var dates models.DatesIndex

		urls := map[any]string{
			&artists:   config.ArtistsURL,
			&relations: config.RelationsURL,
			&locations: config.LocationURL,
			&dates:     config.DatesURL,
		}

		var wg sync.WaitGroup
		for a, u := range urls {
			wg.Add(1)
			go func(an *any, url string) {
				biblio.Help.Fetch(url, an)
				wg.Done()
				}(&a, u)
		}
		wg.Wait()

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
		var err error
		if models.AtStartingServer {
			if err = biblio.Help.RenderPageInsideBuffer(config.Pages+"index", nil, &models.HomeTemplate); err != nil {
				return err
			}
			if err = biblio.Help.RenderPageInsideBuffer(config.Pages+"artists", struct{ Artists []models.ArtistAllData }{Artists: models.ArtistsFullData}, &models.ArtistsTemplate); err != nil {
				return err
			}
			models.AtStartingServer = false
		} else {
			if err = biblio.Help.RenderPageInsideBuffer(config.Pages+"artists", struct{ Artists []models.ArtistAllData }{Artists: models.ArtistsFullData}, &models.ArtistsTemplate); err != nil {
				return err
			}
		}

		models.LastTimeFetch = time.Now()
	}
	return nil
}
