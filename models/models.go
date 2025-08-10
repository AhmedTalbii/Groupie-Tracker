package models

import (
	"bytes"
	"groupie-tracker/config"
	"sync"
	"time"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	LocationsUrl string   `json:"locations"`
	DatesUrl     string   `json:"concertDates"`
	RelationsUrl string   `json:"relations"`
}

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type LocationsIndex struct {
	Index []Locations `json:"index"`
}

type DatesIndex struct {
	Index []Dates `json:"index"`
}

type RelationsIndex struct {
	Index []Relations `json:"index"`
}

// all data so that we can acces to all the strcucts easly
type AllData struct {
	Artists   []Artist
	Locations LocationsIndex
	Dates     DatesIndex
	Relations RelationsIndex
}


// this stuct contain all the data for etch artist
type ArtistAllData struct {
	Id           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Relations    map[string][]string
	Locations    []string
	Dates        []string
	ArtistURL    string
	RelationsURL string
	LocationsURL string
	DatesURL     string
}

var ArtistsFullData []ArtistAllData

// Templates
var (
	HomeTemplate    bytes.Buffer
	ArtistsTemplate bytes.Buffer
)

// ttl 
var (
	Mu            sync.Mutex
	LastTimeFetch time.Time = time.Now().Add(-config.TimeToRefreshData)
	AtStartingServer = true
)
