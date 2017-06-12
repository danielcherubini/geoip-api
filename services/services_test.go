package services

import (
	"fmt"
	"testing"

	"github.com/danmademe/geoip-api/models"
)

func TestDownloadS3Url(t *testing.T) {
	fmt.Println("Testing DownloadS3Url")

	s3Config := models.S3Config{}
	s3Config.Bucket = "tidal-bi-emr"
	s3Config.Key = "/libs/playbacklog/GeoIP2-City.mmdb"
	s3Config.Region = "us-east-1"

	filename := "GeoIP2-City.mmdb"

	err, filePath := DownloadS3Url(s3Config, filename)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println("Test Pass: " + filePath)
}

//
// func TestDownloadUrl(t *testing.T) {
// fmt.Println("Testing DownloadS3Url")
// 	urlString := "http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz"
// 	filename := "GeoLite2-City.tar.gz"
//
// 	err, filePath := DownloadUrl(urlString, filename)
// 	if err != nil {
// 		fmt.Println(err)
// 		t.Fail()
// 	}
//
// fmt.Println("Test Pass: " + filePath)
// }
