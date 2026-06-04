package main

import (
	"log"
	"pitsinkr/internals/dns"
	"pitsinkr/internals/storage"
	"pitsinkr/internals/web"
	"server/api"
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
