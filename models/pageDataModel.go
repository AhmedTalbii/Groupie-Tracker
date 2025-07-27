package models

import (
	"sync"
	"time"

	"groupietracker/config"
)

// artist data all the data from the 4 apis in one place
type ArtistData struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	LocationsUrl string
	DatesUrl     string
	RelationsUrl string
	Locations    []string
	Dates        []string
	Relations    map[string][]string
}

// struct to make the artists in one place so we can put  it in the page
type PageData struct {
	Artists []ArtistData
}

// mutex + data update + fetch
var (
	DataFetched    []ArtistData
	LastUpdateTime time.Time
	Ttl            = config.Ttl * time.Minute
	Mu             sync.Mutex
	InitFetch      = true
)
