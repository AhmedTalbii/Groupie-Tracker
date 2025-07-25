package models

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

type LocationsIndex struct {
	Index []Locations `json:"index"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DatesIndex struct {
	Index []Dates `json:"index"`
}

type Relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationsIndex struct {
	Index []Relations `json:"index"`
}

type AllData struct {
	Artists   []Artist
	Locations LocationsIndex
	Dates     DatesIndex
	Relations RelationsIndex
}
