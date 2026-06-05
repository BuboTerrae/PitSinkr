package dns

import (
	"fmt"
	"log"
	"net"
	"pitsinkr/internal/storage"
	"time"

	"github.com/miekg/dns"
)

type DNSLog struct {
	ClientIP string
	Domain   string
	Type     string
	Time     time.Time
}

func handleDNS(w dns.ResponseWriter, r *dns.Msg) {
	resp, err := Resolve(r)
	if err != nil {
		fmt.Printf("Resolve Error: %v\n", err)
	}

	domain := r.Question[0].Name
	qtype := dns.TypeToString[r.Question[0].Qtype]
	host, port, splitErr := net.SplitHostPort(w.RemoteAddr().String())

	if splitErr != nil {
		return
	}
	fmt.Printf("[%s] %s | Port: %s | Type: %s\n", host, domain, port, qtype)
	entry := storage.DNSLog{ClientIP: host, Domain: domain, Type: qtype}

	db, err := storage.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	err = storage.InsertLog(db, entry)
	if err != nil {
		fmt.Printf("Insert Error: %v\n", err)
	}

	w.WriteMsg(resp)
}
