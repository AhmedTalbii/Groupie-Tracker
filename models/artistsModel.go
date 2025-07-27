package models

// thoes strcucts so that we can unmarchal the data from apis to structured data
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

// thoes struct cause somme apis have index witch is array thats whhy we make them 
type LocationsIndex struct {
	Index []Locations `json:"index"`
}

type DatesIndex struct {
	Index []Dates `json:"index"`
}

type RelationsIndex struct {
	Index []Relations `json:"index"`
}

// all data so that we can acces to all the data easly 
type AllData struct {
	Artists   []Artist
	Locations LocationsIndex
	Dates     DatesIndex
	Relations RelationsIndex
}
