# Groupie Trackers

Groupie Trackers is a web application built in Go that consumes a provided API to display information about artists. The API includes:

- **Artists**: Names, images, start year, first album date, members.
- **Locations**: Concert locations.
- **Dates**: Concert dates.
- **Relation**: Links artists, dates, and locations.

## Link to the hosted website
[Groupie Tracker](https://gtia.up.railway.app/)

## Authors
- Aboudou Ilyass
- Talbi Ahmed

## Features
- User-friendly display using cards, tables, or lists.
- Client-server communication with a custom event/action that triggers a request to the server.
- Fully functional backend in Go.
- Error handling to ensure stability.

## Requirements
- Backend: **Go**
- Follow good coding practices.

## How to Run
```bash
git clone https://learn.zone01oujda.ma/git/ahtalbi/groupie-tracker.git
cd groupie-tracker
go run .
```

## Notes
This project was developed as part of the **Zone01 Cursus**.


## structure
```
Groupie Tracher
|
├── README.md
├── biblio
    ├── biblio.go
    └── pages
    │   └── error.html
├── config
    └── config.go
├── controllers
    ├── fetchers
    │   └── init.go
    └── handlers
    │   ├── artists.go
    │   ├── home.go
    │   └── static.go
├── go.mod
├── main.go
├── models
    └── models.go
├── routes
    └── router.go
├── server
    └── server.go
├── statics
    ├── assets
    │   ├── Logo.png
    │   ├── mockup.png
    │   └── s.png
    └── css
    │   ├── artist.css
    │   ├── artists.css
    │   └── index.css
