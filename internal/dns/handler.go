package dns

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"pitsinkr/internal/storage"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/miekg/dns"
)

type DNSLog struct {
	ClientIP string
	Domain   string
	Type     string
	Time     time.Time
}

var (
	blocklist  = make(map[string]bool)
	blockMutex sync.RWMutex
)

func handleDNS(w dns.ResponseWriter, r *dns.Msg) {
	if len(r.Question) == 0 {
		return
	}

	q := r.Question[0]
	domain := q.Name
	qtype := dns.TypeToString[q.Qtype]
	host, port, _ := net.SplitHostPort(w.RemoteAddr().String())

	blocked, reason := isBlocked(domain)
	logEntry := storage.DNSLog{ClientIP: host, Domain: domain, Type: qtype, Blocked: blocked, Reason: reason, Timestamp: time.Now()}

	db, err := storage.OpenDB()
	if err != nil {
		log.Fatal(err)
	}
	err = storage.InsertLog(db, logEntry)
	if err != nil {
		fmt.Printf("Insert Error: %v\n", err)
	}

	fmt.Printf("[%s]:%s %s %s %s\n", host, port, domain, qtype, map[bool]string{true: "BLOCKED", false: "ALLOWED"}[blocked])

	if blocked {
		resp := &dns.Msg{}
		resp.SetReply(r)
		resp.SetRcode(r, dns.RcodeSuccess)

		switch q.Qtype {
		case dns.TypeA:
			resp.Answer = append(resp.Answer, &dns.A{
				Hdr: dns.RR_Header{Name: domain, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 86400},
				A:   net.ParseIP("0.0.0.0"),
			})
		case dns.TypeAAAA:
			resp.Answer = append(resp.Answer, &dns.AAAA{
				Hdr:  dns.RR_Header{Name: domain, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: 86400},
				AAAA: net.ParseIP("::"),
			})
		default:
			resp.SetRcode(r, dns.RcodeNameError)
		}

		w.WriteMsg(resp)
		return
	}

	resp, err := Resolve(r)
	if err != nil {
		fmt.Printf("Resolve Error: %v\n", err)
		resp = new(dns.Msg)
		resp.SetRcode(r, dns.RcodeServerFailure)
	}

	w.WriteMsg(resp)
}

func loadBlocklist(paths []string) error {
	blockMutex.Lock()
	defer blockMutex.Unlock()

	clear(blocklist)

	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			continue
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if line == "" || strings.HasPrefix(line, "#") {
				continue
			}
			dom := strings.ToLower(strings.Trim(line, "."))
			blocklist[dom] = true
		}
	}
	return nil
}

func isBlocked(domain string) (bool, string) {
	norm := normalizeDomain(domain)

	blockMutex.RLock()
	defer blockMutex.RUnlock()

	if blocklist[norm] {
		return true, "blocklist"
	}

	parts := strings.Split(norm, ".")
	for i := 1; i < len(parts); i++ {
		suffix := strings.Join(parts[i:], ".")
		if blocklist[suffix] {
			return true, "blocklist"
		}
	}

	return false, ""
}

func normalizeDomain(domain string) string {
	d := strings.ToLower(strings.TrimSpace(domain))
	d = strings.Trim(d, ".")
	return d
}

func AddDomain(domain string) {
	norm := normalizeDomain(domain)
	if norm == "" {
		return
	}

	blockMutex.Lock()
	defer blockMutex.Unlock()

	blocklist[norm] = true
	fmt.Printf("Blocked: %s\n", domain)
}

func RemoveDomain(domain string) {
	norm := normalizeDomain(domain)
	if norm == "" {
		return
	}

	blockMutex.Lock()
	defer blockMutex.Unlock()

	delete(blocklist, norm)
	fmt.Printf("Unblocked: %s\n", norm)
}

func GetBlocklist() []string {
	blockMutex.RLock()
	defer blockMutex.RUnlock()

	list := make([]string, 0, len(blocklist))
	for d := range blocklist {
		list = append(list, d)
	}
	sort.Strings(list)
	return list
}
