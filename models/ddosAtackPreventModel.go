package models

import "time"

type Client struct {
	Requests int
	LastTime time.Time
	Warning  bool
	Banned   bool
}

// clients
var (
	Clients map[string]*Client
)
