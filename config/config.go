package config

import "time"

// Port
const (
	Port = ":8080"
)

// Paths
const (
	Pages = "views/"
)

// Api's
const (
	ArtistsURL   = "https://groupietrackers.herokuapp.com/api/artists"
	LocationURL  = "https://groupietrackers.herokuapp.com/api/locations"
	RelationsURL = "https://groupietrackers.herokuapp.com/api/relation"
	DatesURL     = "https://groupietrackers.herokuapp.com/api/dates"
)

// Ttl
const (
	TimeToRefreshData = 10 * time.Minute // in minuts
)
