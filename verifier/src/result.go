package main

import "fmt"

type Result struct {
	TcpTimeout   uint
	TlsHandshake uint
	Unexpected   uint
	Success      uint
}

func (r *Result) String() string {
	return fmt.Sprintf("%d,%d,%d,%d", r.TcpTimeout, r.TlsHandshake, r.Unexpected, r.Success)
}
