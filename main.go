package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/danmademe/geoip-api/utils"
	"github.com/gorilla/mux"
)

//CheckIPRoute is the route for checking IP
//Returns with ip address
func CheckIPRoute(w http.ResponseWriter, r *http.Request) {
	//Setup struct
	responseObject := utils.GetLocale(r)

	//Marshal responseObject to JsonResponse
	jsonByteArray, _ := json.Marshal(responseObject)
	jsonString := string(jsonByteArray)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", jsonString)
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

func main() {
	log.Fatal(Setup().ListenAndServe())
}
