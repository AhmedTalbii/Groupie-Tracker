package models

type ArtistData struct {
	Id           int
	Image        string
	Name         string
	CreationDate int
	FirstAlbum   string
	Members      []string
	Concerts     map[string][]string
	Locations    []string
	Dates        []string
	Sources      []string
}

type PageData struct {
	Artists []ArtistData
}

var (
	DataFetched []ArtistData
)
