package utils

import (
	"fmt"
	"testing"
)

func TestDownloadUrl(t *testing.T) {
	urlString := "https://s3.amazonaws.com/tidal-bi-emr/libs/playbacklog/GeoIP2-City.mmdb"
	filename := "GeoIP2-City.mmdb"

	err, filePath := downloadUrl(urlString, filename)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(filePath)
}
