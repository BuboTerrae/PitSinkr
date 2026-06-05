package main

import (
	"log"
	"pitsinkr/internal/dns"
	"pitsinkr/internal/storage"
	"pitsinkr/internal/web"
	"pitsinkr/cmd/server/api"
)

func main() {
	db, err := storage.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	storage.InitDB(db)

	go dns.Start()
	go api.Start(db)

	web.StartWebServer()
}
