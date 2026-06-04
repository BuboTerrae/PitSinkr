package dns

import (
	"github.com/miekg/dns"
)

func Resolve(req *dns.Msg) (*dns.Msg, error) {
	client := new(dns.Client)

	req, _, err := client.Exchange(req, "1.1.1.1:53")

	return req, err
}
