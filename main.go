package main

import (
	"groupietracker/server"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1  {
		log.Fatal("No arguments needed ...")
	}
	server.StartServer()
}