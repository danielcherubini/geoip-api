package utils

import (
	"net"
	"net/http"
	"strings"
)

//GetIP takes a request and returns a string
func GetIP(r *http.Request) net.IP {
	remoteIP := r.RemoteAddr
	ipString := ""

	if (remoteIP == "") || (strings.Contains(remoteIP, "127.0.0.1")) {
		forwaredForIP := r.Header.Get("X-Forwarded-For")
		ipString = forwaredForIP
	} else {
		ipString = remoteIP
	}

	return net.ParseIP(ipString)
}
