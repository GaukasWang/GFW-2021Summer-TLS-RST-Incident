package main

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	tls "github.com/refraction-networking/utls"
	"golang.org/x/net/http2"
)

var (
	dialTimeout = time.Duration(15) * time.Second
)

func httpGetOverConn(conn net.Conn, alpn string, hostname string) (*http.Response, error) {
	req := &http.Request{
		Method: "GET",
		URL:    &url.URL{Host: "www." + hostname + "/"},
		Header: make(http.Header),
		Host:   "www." + hostname,
	}

	switch alpn {
	case "h2":
		req.Proto = "HTTP/2.0"
		req.ProtoMajor = 2
		req.ProtoMinor = 0

		tr := http2.Transport{}
		cConn, err := tr.NewClientConn(conn)
		if err != nil {
			return nil, err
		}
		return cConn.RoundTrip(req)
	case "http/1.1", "":
		req.Proto = "HTTP/1.1"
		req.ProtoMajor = 1
		req.ProtoMinor = 1

		err := req.Write(conn)
		if err != nil {
			return nil, err
		}
		return http.ReadResponse(bufio.NewReader(conn), req)
	default:
		return nil, errALPNUnsupported
	}
}

func TestHTTPSHandshake(hostname string, addr string, clientHelloSpec tls.ClientHelloSpec) (*http.Response, error) {
	tcpConn, err := net.DialTimeout("tcp", addr, dialTimeout)
	if err != nil {
		fmt.Printf("net.Dial() failed: %+v\n", err)
		return nil, errDialTimeout
	}

	config := tls.Config{ServerName: hostname}
	// This fingerprint includes feature(s), not fully supported by TLS.
	// uTLS client with this fingerprint will only be able to to talk to servers,
	// that also do not support those features.
	tlsConn := tls.UClient(tcpConn, &config, tls.HelloCustom)
	err = tlsConn.ApplyPreset(&clientHelloSpec)
	if err != nil {
		return nil, errApplyPreset
	}

	// n, err = tlsConn.Write("Hello, World!")
	err = tlsConn.Handshake()
	if err != nil {
		return nil, errHandshake
	}

	return nil, nil
	// return httpGetOverConn(tlsConn, tlsConn.HandshakeState.ServerHello.AlpnProtocol, hostname)
}
