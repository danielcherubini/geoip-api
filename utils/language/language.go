package language

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/danmademe/geoip-api/models"
)

//LoadLanguages loads languages to a languages model
func LoadLanguages(langFile string) []models.Language {
	file, err := ioutil.ReadFile(langFile)
	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}
	var language []models.Language
	json.Unmarshal(file, &language)
	return language
}

//GetLanguage takes a string and returns a language model
func GetLanguage(country string) models.Language {
	language := models.Language{}

	for i := 0; i < len(models.Languages); i++ {
		if strings.ToLower(models.Languages[i].Country) == strings.ToLower(country) {
			language = models.Languages[i]
			return language
		}

	}
	return language
}
