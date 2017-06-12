package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/danmademe/geoip-api/models"
	"github.com/danmademe/geoip-api/utils"
	"github.com/gorilla/mux"
)

var currentHash = ""
var ready = false
var mmdb string
var gzdb string
var dburl string
var s3bucket string
var s3key string
var s3region string
var langFile string

//CheckIPRoute is the route for checking IP
//Returns with ip address
func CheckIPRoute(w http.ResponseWriter, r *http.Request) {
	//Check if ip query is missing
	if r.URL.Query().Get("ip") == "" {
		fmt.Fprintf(w, "{\"error\": \"You didn't pass an IP\"}")
	} else {
		//Setup struct
		err, responseObject := utils.GetLocale(r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
		} else {

			//Marshal responseObject to JsonResponse
			jsonByteArray, _ := json.Marshal(responseObject)
			jsonString := string(jsonByteArray)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s", jsonString)
		}
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
	fmt.Println("Starting Router")
	r := Router()
	http.Handle("/", r)
	srv := Server(r)
	return srv
}

//Load this gets the config file from flag
// accepts --lang flag
func Load() []models.Language {
	if models.ConfigFile != "" {
		langFile = models.ConfigFile
	} else {
		models.ConfigFile = langFile
	}

	lang := utils.LoadLanguages(langFile)
	return lang
}

func getDatabase(db models.DBLocation) {
	err := utils.GetDatabase(db)
	if err != nil {
		os.Exit(1)
	}
}

func checkS3vars(s3Config models.S3Config) {
	if s3Config.Bucket == "" || s3Config.Key == "" || s3Config.Region == "" {
		fmt.Println("S3 Config incorrect missing params")
		os.Exit(1)
	}
}

func setupVars() models.DBLocation {
	var db models.DBLocation
	flag.StringVar(&langFile, "lang", "", "Location of language.json file (REQUIRED)")
	flag.StringVar(&mmdb, "mmdb", "", "Location of local .mmdb file")
	flag.StringVar(&gzdb, "gzdb", "", "Location of local .gzip file")
	flag.StringVar(&dburl, "dburl", "", "Location of remote mmdb/gzip file")
	flag.StringVar(&s3bucket, "s3bucket", "", "Bucket name on S3")
	flag.StringVar(&s3key, "s3key", "", "Full file path from root on the s3 bucket so if its /foo/bar/qux.jpg then thats the key")
	flag.StringVar(&s3region, "s3region", "", "Which region is the S3 bucket in")

	flag.Parse()

	if langFile == "" {
		fmt.Printf("File error: Missing lang file param pass in via --lang language.json\n")
		os.Exit(1)
	}

	if mmdb != "" {
		db.Type = "mmdb"
		db.Location = mmdb
	} else if gzdb != "" {
		db.Type = "GZDB"
		db.Location = gzdb
	} else if dburl != "" {
		db.Type = "DBURL"
		db.Location = dburl
	} else if s3bucket != "" {
		db.Type = "S3DB"
		db.S3Config = models.S3Config{}
		db.S3Config.Bucket = s3bucket
		db.S3Config.Key = s3key
		db.S3Config.Region = s3region
		checkS3vars(db.S3Config)
	} else {
		db.Type = "DEFAULT"
		db.Location = "http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz"
	}

	return db
}

func main() {
	db := setupVars()
	err := utils.GetDatabase(db)
	if err != nil {
		os.Exit(1)
	} else {
		ticker := time.NewTicker(24 * time.Hour)
		quit := make(chan struct{})
		go func() {
			for {
				select {
				case <-ticker.C:
					err := utils.GetDatabase(db)
					if err != nil {
						os.Exit(1)
					}
				case <-quit:
					ticker.Stop()
					return
				}
			}
		}()
		models.Languages = Load()
		fmt.Println("Loaded Languages")
		log.Fatal(Setup().ListenAndServe())
	}

}
