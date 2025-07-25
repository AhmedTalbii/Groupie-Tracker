package fetchers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func FetchData[T any](url string) (T, error) {
	// get the response from url
	res, errorGetingTheData := http.Get(url)
	var zero T // we declare this so if we have an error cause
	if errorGetingTheData != nil {
		return zero, errorGetingTheData
	}
	defer res.Body.Close()

	var data T
	dataBytes, errorReadingTheResBody := io.ReadAll(res.Body)
	if errorReadingTheResBody != nil {
		return zero, errorReadingTheResBody
	}
	errorUnmarchal := json.Unmarshal(dataBytes, &data)
	if errorUnmarchal != nil {
		return zero, errorUnmarchal
	}
	return data, nil
}

func MustFetch[T any](url string) T {
	data, err := FetchData[T](url)
	if err != nil {
		log.Fatal(err)
	}
	return data
}