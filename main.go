package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danmademe/geoip-api/models"
	"github.com/danmademe/geoip-api/utils"
	"github.com/gorilla/mux"
)

//CheckIPRoute is the route for checking IP
//Returns with ip address
func CheckIPRoute(w http.ResponseWriter, r *http.Request) {
	//Check if ip query is missing
	if r.URL.Query().Get("ip") == "" {
		fmt.Fprintf(w, "{\"error\": \"You didn't pass an IP\"}")
	} else {
		//Setup struct
		responseObject := utils.GetLocale(r)

		//Marshal responseObject to JsonResponse
		jsonByteArray, _ := json.Marshal(responseObject)
		jsonString := string(jsonByteArray)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", jsonString)
	}
}

//Router dfdfd
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", CheckIPRoute)

	return r
}

//Server setup for server
func Server(r *mux.Router) *http.Server {
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:45000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	return srv
}

//Setup sets server up for logfatal
func Setup() *http.Server {

	r := Router()
	http.Handle("/", r)
	srv := Server(r)
	return srv
}

//Load this gets the config file from flag
// accepts --lang flag
func Load() []models.Language {
	var langFile string
	flag.StringVar(&langFile, "lang", "./languages.json", "a string var")
	flag.Parse()

	models.ConfigFile = langFile
	lang := utils.LoadLanguages(langFile)
	return lang
}

func main() {
	models.Languages = Load()
	log.Fatal(Setup().ListenAndServe())
}
