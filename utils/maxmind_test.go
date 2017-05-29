package utils

import (
	"testing"

	"github.com/danmademe/geoip-api/models"
)

func TestGetDatabase(t *testing.T) {
	db := models.DBLocation{
		Location: "http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz",
		Type:     "DBURL",
	}
	// db := models.DBLocation{
	// 	Location: "https://s3.amazonaws.com/tidal-bi-emr/libs/playbacklog/GeoIP2-City.mmdb",
	// 	Type:     "DBURL",
	// }
	// db := models.DBLocation{
	// 	Location: "../GeoLite2-City.tar.gz",
	// 	Type:     "GZDB",
	// }
	// db := models.DBLocation{
	// 	Location: "../GeoLite2-City_20170502/GeoLite2-City.mmdb",
	// 	Type:     "MMDB",
	// }
	GetDatabase(db)
}
