package fetchers

import (
	"encoding/json"
	"net/http"
)

func Fetch(url string, result interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return err
}

func MustFetch(url string, result interface{}) {
	err := Fetch(url, result)
	if err != nil {
		// 500 internal server
		return
	}
}
