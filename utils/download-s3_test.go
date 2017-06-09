package utils

import (
	"fmt"
	"testing"

	"github.com/danmademe/geoip-api/models"
)

func TestDownloadS3Url(t *testing.T) {
	s3Config := models.S3Config{}
	s3Config.Bucket = "tidal-bi-emr"
	s3Config.Key = "/libs/playbacklog/GeoIP2-City.mmdb"
	s3Config.Region = "us-east-1"

	filename := "GeoIP2-City.mmdb"

	err, filePath := downloadS3Url(s3Config, filename)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	fmt.Println(filePath)
}
