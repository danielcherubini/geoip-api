package utils

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

func getLanguage(country string) models.Language {
	language := models.Language{}
	languages := LoadLanguages("./languages.json")

	for i := 0; i < len(languages); i++ {
		if strings.ToLower(languages[i].Country) == strings.ToLower(country) {
			language = languages[i]
			return language
		}

	}
	return language
}
