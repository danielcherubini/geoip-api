package utils

import (
	"errors"
	"log"
	"net"
	"net/http"

	"github.com/danmademe/geoip-api/models"
	"github.com/danmademe/geoip-api/utils/ip"
	"github.com/danmademe/geoip-api/utils/language"
	geoip2 "github.com/oschwald/geoip2-golang"
)

//GetCountry takes an ipString and returns a country
func GetCountry(ip net.IP) *geoip2.Country {
	db, err := geoip2.Open("geo.mmdb")
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
func GetLocale(r *http.Request) (error, models.Response) {
	locale := &models.Response{}

	ipString := r.URL.Query().Get("ip")
	ip := ip.GetIP(ipString)
	if ip == nil {
		err := errors.New("Invalid IP")
		return err, *locale
	} else {
		countryCode := GetCountry(ip).Country.IsoCode

		language := language.GetLanguage(countryCode).Language
		locale.IPAddress = ip.String()
		locale.CountryCode = countryCode
		locale.Language = language
		locale.IsoString = language + "-" + countryCode
		return nil, *locale
	}

}
