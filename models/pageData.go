package models

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

type PageData struct {
	Artists []ArtistData
}

var DataFetched []ArtistData
