package fetchers

import (
	"fmt"
	"time"

	"groupietracker/config"
	"groupietracker/models"
)

func InitFetch() {
	if models.InitFetch {
		fmt.Println("🔄 Fetching artist data... please wait")
	}

	var pageData []models.ArtistData

	// fetch all the data
	Artists := MustFetch[[]models.Artist](config.API_ARTISTS_URL)
	Locations := MustFetch[models.LocationsIndex](config.API_LOCATIONS_URL)
	Dates := MustFetch[models.DatesIndex](config.API_DATES_URL)
	Relations := MustFetch[models.RelationsIndex](config.API_RELATIONS_URL)

	// struct to optimize calling
	AllData := &models.AllData{
		Artists:   Artists,
		Locations: Locations,
		Dates:     Dates,
		Relations: Relations,
	}

	// for loop for eatch artist
	for i := 0; i < len(Artists); i++ {
		Artist := &models.ArtistData{
			Id:           AllData.Artists[i].Id,
			Image:        AllData.Artists[i].Image,
			Name:         AllData.Artists[i].Name,
			Members:      AllData.Artists[i].Members,
			CreationDate: AllData.Artists[i].CreationDate,
			FirstAlbum:   AllData.Artists[i].FirstAlbum,
			LocationsUrl: AllData.Artists[i].LocationsUrl,
			DatesUrl:     AllData.Artists[i].DatesUrl,
			RelationsUrl: AllData.Artists[i].RelationsUrl,
			Locations:    AllData.Locations.Index[i].Locations,
			Dates:        AllData.Dates.Index[i].Dates,
			Relations:    AllData.Relations.Index[i].DatesLocations,
		}
		pageData = append(pageData, *Artist)
	}

	// put in the global variable the data
	models.DataFetched = pageData

	// update the last time that we fetch 
	models.LastUpdateTime = time.Now()
	if models.InitFetch {
		fmt.Println("✅ Done fetching artists data you can acces the link (:")
		models.InitFetch = false
	}
}
