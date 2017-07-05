package ip

import (
	"net"
	"strings"
)

//GetIP takes a string and parses the ip with the string
func GetIP(ipString string) net.IP {
	ip := strings.Split(ipString, ", ")
	return net.ParseIP(ip[0])
}
