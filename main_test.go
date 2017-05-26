package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetup(t *testing.T) {
	srv := Setup()
	if srv == nil {
		t.Fail()
	}
}

func TestCheckIpRoute(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	req.Header.Set("X-Forwarded-For", "193.215.2.26")
	handler := http.HandlerFunc(CheckIPRoute)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `{"ip_address":"193.215.2.26","country_code":"NO","language":"en","iso":"en-NO"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
