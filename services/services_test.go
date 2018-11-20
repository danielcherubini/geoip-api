package services

import (
	"fmt"
	"testing"
)

func TestDownloadUrl(t *testing.T) {
	fmt.Println("Testing DownloadS3Url")
	urlString := "http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz"
	filename := "GeoLite2-City.tar.gz"

	err, filePath := DownloadUrl(urlString, filename)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println("Test Pass: " + filePath)
}
