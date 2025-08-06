package models

// full artists data
type FullArtistsData struct {
	Artist    *Artist
	Locations *Locations
	Relations *Relations
	Dates     *Dates
}

var Artists []Artist

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
type Locations struct {
	Index IndL `json:"index"`
}

type IndL []struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

// relations
type Relations struct {
	Index IndR
}

type IndR []struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// dates
type Dates struct {
	Index IndD
}

type IndD []struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
