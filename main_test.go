package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danielcherubini/geoip-api/models"
	"github.com/danielcherubini/geoip-api/utils/language"
)

func TestSetup(t *testing.T) {
	srv := Setup()
	if srv == nil {
		t.Fail()
	}
}

func TestGetDatabase(t *testing.T) {
	db := models.DBLocation{
		Location: "http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz",
		Type:     "DBURL",
	}
	getDatabase(db)
}

func TestLoad(t *testing.T) {
	models.ConfigFile = "./languages.json"
	languages := Load()

	if languages.Languages[1].Country == "" {
		t.Fail()
	}
}

func TestCheckIpRoute(t *testing.T) {
	//load dummy data
	models.LanguageConfig = language.LoadLanguages("./languages.json")

	req, err := http.NewRequest("GET", "/?ip=193.215.2.26", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(CheckIPRoute)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"ip_address":"193.215.2.26","country_code":"NO","language":"no","iso":"no_NO"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCheckIpRouteLocalhost(t *testing.T) {
	//load dummy data
	models.LanguageConfig = language.LoadLanguages("./languages.json")

	req, err := http.NewRequest("GET", "/?ip=127.0.0.1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(CheckIPRoute)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"ip_address":"127.0.0.1","country_code":"NO","language":"no","iso":"no_NO"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
