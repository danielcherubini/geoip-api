package utils

import (
	"fmt"
	"testing"
)

func TestGetLanguage(t *testing.T) {
	language := getLanguage("NO")

	if language.Language != "en" {
		t.Fail()
	}

	fmt.Println(language)
}
