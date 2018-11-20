package utils

import (
	"testing"

	"github.com/danmademe/geoip-api/models"
	"github.com/danmademe/geoip-api/utils/ip"
	"github.com/danmademe/geoip-api/utils/language"
)

func TestGetCountry(t *testing.T) {
	singleIP := "193.215.2.26"
	ip := ip.GetIP(singleIP)
	countryCode := GetCountry(ip)
	models.LanguageConfig = language.LoadLanguages("../languages.json")
	language := language.GetLanguage(countryCode)

	if countryCode != "NO" {
		t.Errorf("Wrong countrycode wanted 'NO' got %v", countryCode)
	}
	if language.Language != "no" {
		t.Errorf("Wrong language wanted 'no' got %v", language.Language)
	}
}
