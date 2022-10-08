package main

import (
	"fmt"
	"net"
	"time"

	tls "github.com/refraction-networking/utls"
)

var (
	dialTimeout time.Duration
	iter        uint
	slp         time.Duration
)

func TLSHandshake(SNI, addr string) {
	mapResults := make(map[string]Result)
	// Iterate through the client hello specs
	for chName, clientHelloGen := range clientHellos {
		result := Result{}
		for i := uint(0); i < iter; i++ {
			time.Sleep(slp)
			tcpConn, err := net.DialTimeout("tcp", addr, dialTimeout)
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					result.TcpTimeout++
				} else {
					result.Unexpected++
				}
				continue
			}

			config := tls.Config{ServerName: SNI}
			// This fingerprint includes feature(s), not fully supported by TLS.
			// uTLS client with this fingerprint will only be able to to talk to servers,
			// that also do not support those features.
			tlsConn := tls.UClient(tcpConn, &config, tls.HelloCustom)
			clientHello := clientHelloGen()

			err = tlsConn.ApplyPreset(&clientHello)
			if err != nil {
				result.Unexpected++
				continue // ignore
			}

			err = tlsConn.Handshake()
			if err != nil {
				result.TlsHandshake++
				continue
			}

			result.Success++
		}
		// push the result to the map
		mapResults[chName] = result
	}

	// Print the results
	resultStr := fmt.Sprintf("==== %s, %s ====\n", SNI, addr)
	for chName, result := range mapResults {
		resultStr += fmt.Sprintf("%s: %s\n", chName, result.String())
	}

	resultStr += "\n"

	fmt.Println(resultStr)
}
