package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"pitsinkr/cmd/server/api"
	"pitsinkr/internal/dns"
	"pitsinkr/internal/storage"
	"pitsinkr/internal/web"
	"syscall"
)

const (
	DNSListenAddr  = ":53"
	SNIListenAddr  = ":8443" // redirected from 443
	HTTPListenAddr = ":8080" // redirected from 80
	DBPath         = "./dns_sinkhole.db"
)

func main() {
	db, err := storage.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	storage.InitDB(db)

	dns.AddDomain("youtube.com")

	go dns.Start()
	go api.Start(db)
	web.StartWebServer()

	fmt.Println("Sinkhole running - DNS:", DNSListenAddr, "| SNI: ", SNIListenAddr, "| HTTP:", HTTPListenAddr)
	fmt.Println("Make sure iptables redirect 80=>8080 and 443=>8443")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("Shutting down...")
}
