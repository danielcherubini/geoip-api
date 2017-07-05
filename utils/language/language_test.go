package language

import (
	"fmt"
	"testing"

	"github.com/danmademe/geoip-api/models"
)

func TestGetLanguage(t *testing.T) {
	models.Languages = LoadLanguages("../../languages.json")
	language := GetLanguage("NO")

	if language.Language != "en" {
		t.Fail()
	}

	fmt.Println(language)
}
