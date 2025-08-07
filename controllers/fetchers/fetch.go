package fetchers

import (
	"encoding/json"
	"net/http"
)

// sends HTTP GET request to URL,
// decodes the JSON response into 'result',
// and returns any error if exist.
func Fetch(url string, result interface{}, w http.ResponseWriter) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&result)
	return err
}

// calls Fetch and silently returns on error,
func MustFetch(url string, result interface{}, w http.ResponseWriter) {
	err := Fetch(url, &result, w)
	if err != nil {
		return
	}
}
