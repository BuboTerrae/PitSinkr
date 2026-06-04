package dns

import "github.com/miekg/dns"

func Start() {
	dns.HandleFunc(".", handleDNS)

	server := &dns.Server{Addr: ":53", Net: "udp"}

	server.ListenAndServe()
}
