package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//GetDatabase gets the database
func GetDatabase() {
	resp, err := http.Get("http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz.md5")
	if err != nil {
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
