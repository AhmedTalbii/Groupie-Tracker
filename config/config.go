package config

// Port
const (
	Port = ":8000"
)

// APIS
const (
	API_ARTISTS_URL   = "https://groupietrackers.herokuapp.com/api/artists"
	API_LOCATIONS_URL = "https://groupietrackers.herokuapp.com/api/locations"
	API_DATES_URL     = "https://groupietrackers.herokuapp.com/api/dates"
	API_RELATIONS_URL = "https://groupietrackers.herokuapp.com/api/relation"
)

// Paths
const (
	Layout     = "views/"
	Pages      = "views/pages/"
	Sections   = "views/sections/"
	Components = "views/components/"
)

// Ttl
const (
	Ttl = 10 // in minuts 
)

// DDOS Atack
const (
	RequestsNumbertToGetWarningOrBan = 20
	TimeBetweenEatchReq = 4 // in secends
	BannTime = 1 // in hours
)