package fetchers

import (
	"encoding/json"
	"groupietracker/models"
	"io"
	"net/http"
)

func FetchData(URL string) ([]models.Artist, error) {
	// get data as res from url 
	res, errGettingData := http.Get(URL)
	if errGettingData != nil {
		return nil, errGettingData
	}
	// read the body of the res using io 
	data, errIoInFetch := io.ReadAll(res.Body)
	if errIoInFetch != nil {
		return nil, errIoInFetch
	}
	// define artists 
	var artists = []models.Artist{}

	// unmarchal data from json to array of artists witch we declare in models artists
	errUnmarchal := json.Unmarshal(data, &artists)
	if errUnmarchal != nil {
		return nil, errUnmarchal
	}

	return artists, nil
}