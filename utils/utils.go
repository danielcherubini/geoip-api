package utils

import (
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/danmademe/geoip-api/models"
	geoip2 "github.com/oschwald/geoip2-golang"
)

func getIP(r *http.Request) net.IP {
	remoteIP := r.RemoteAddr
	ipString := ""

	if (remoteIP == "") || (strings.Contains(remoteIP, "127.0.0.1")) {
		ipString = r.URL.Query().Get("ip")
	} else {
		ipString = remoteIP
	}

	return net.ParseIP(ipString)
}

//GetCountry takes an ipString and returns a country
func GetCountry(ip net.IP) *geoip2.Country {
	db, err := geoip2.Open("GeoLite2-Country.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// If you are using strings that may be invalid, check that ip is not nil
	record, err := db.Country(ip)
	if err != nil {
		log.Fatal(err)
	}
	return record
}

//GetLocale takes a country_code and returns object
func GetLocale(r *http.Request) models.Response {
	locale := &models.Response{}

	ip := getIP(r)
	countryCode := GetCountry(ip).Country.IsoCode

	language := getLanguage(countryCode).Language
	locale.IPAddress = ip.String()
	locale.CountryCode = countryCode
	locale.Language = language
	locale.IsoString = language + "-" + countryCode
	return *locale
}
