package language

import (
	"fmt"
	"testing"

	"github.com/danielcherubini/geoip-api/models"
)

func TestGetLanguage(t *testing.T) {
	models.LanguageConfig = LoadLanguages("../../languages.json")
	language := GetLanguage("NO")

	if language.Language != "no" {
		t.Fail()
	}

	fmt.Println(language)
}

func TestGetLanguageOther(t *testing.T) {
	models.LanguageConfig = LoadLanguages("../../languages.json")
	language := GetLanguage("IS")

	if language.Language != "en" {
		t.Fail()
	}

	fmt.Println(language)
}
