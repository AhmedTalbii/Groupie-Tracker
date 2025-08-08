package models

import (
	"html/template"
	"sync"
)

// full artists data
type FullArtistsData struct {
	Artist    *Artist
	Locations *Locations
	Relations *Relations
	Dates     *Dates
}

var (
	Artists []Artist
	Mu      sync.Mutex
)

// artists
type Artist struct {
	Id           int      `json:"id"`
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
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

// relations
type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// dates
type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

var Templat *template.Template

// error page
