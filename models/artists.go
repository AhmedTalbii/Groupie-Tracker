package models

// full artists data
type FullArtistsData struct {
	Artist
	LOCATIONS
	Relations
}

// artists
type Artist struct {
	ID           string   `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// locations
type LOCATIONS struct {
	Index Ind `json:"index"`
}

type Ind []struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

// relations
type Relations struct {
	Index Index
}

type Index []struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
