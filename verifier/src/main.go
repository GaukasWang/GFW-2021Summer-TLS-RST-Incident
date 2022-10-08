package main

import (
	"encoding/csv"
	"flag"
	"log"
	"net"
	"os"
	"sync"
	"time"

	tls "github.com/refraction-networking/utls"
)

type Server struct {
	Addr string
	SNI  string // optional
}

var (
	servers      []*Server                             = []*Server{}
	clientHellos map[string]func() tls.ClientHelloSpec = map[string]func() tls.ClientHelloSpec{
		"XrayCore":  XrayCore,
		"TrojanQt5": TrojanQt5,
		"cURL":      cURL,
		"Chrome106": Chrome106,
	}
)

func main() {
	timeout := flag.Duration("timeout", 5*time.Second, "Duration to wait for communication response")
	scanList := flag.String("list", "list.csv", "Filename of TLS server lists to scan")
	iteration := flag.Uint("iteration", 50, "Number of times to scan each server")
	sleep := flag.Duration("sleep", 5*time.Millisecond, "Duration to sleep between each scan")

	flag.Parse()

	dialTimeout = *timeout

	// Read the list of servers CSV to scan
	f, err := os.Open(*scanList)
	if err != nil {
		log.Fatal("Unable to read input file "+*scanList, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+*scanList, err)
	}

	for _, record := range records {
		server := Server{
			Addr: record[0],
		}
		if len(record) > 1 {
			server.SNI = record[1]
		} else {
			// parse the hostname from the address
			hostname, _, err := net.SplitHostPort(server.Addr)
			if err == nil {
				// if not IP address, use it as SNI
				if net.ParseIP(hostname) == nil {
					server.SNI = hostname
				}
			}
		}
		servers = append(servers, &server)
	}

	iter = *iteration
	slp = *sleep

	wg := sync.WaitGroup{}

	// Scan each server
	for _, server := range servers {
		wg.Add(1)
		go func(sni, addr string) {
			defer wg.Done()
			TLSHandshake(sni, addr)
		}((*server).SNI, (*server).Addr)
	}

	wg.Wait()
}
