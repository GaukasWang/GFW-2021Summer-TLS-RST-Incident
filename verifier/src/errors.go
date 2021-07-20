package main

import "errors"

var (
	errDialTimeout     = errors.New("TCP_DIAL_TIMEOUT")
	errHandshake       = errors.New("HANDSHAKE_FAILED")
	errApplyPreset     = errors.New("APPLYPRESET_FAILED")
	errALPNUnsupported = errors.New("BAD_ALPN")
)
