package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/danmademe/geoip-api/models"
)

func findFile(pattern string) string {

	matches, err := filepath.Glob(pattern)

	if err != nil {
		fmt.Println(err)
	}

	return matches[0]
}

func getDatabase() {
	filename := "GeoLite2-Country.tar.gz"
	url := "http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz"

	output, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error while creating", filename, "-", err)
		return
	}
	defer output.Close()

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}
	defer response.Body.Close()

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", url, "-", err)
		return
	}

	fmt.Println(n, "bytes downloaded.")

	respByte, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("fail to read response data")
		return
	}
	r := bytes.NewReader(respByte)
	Untar("/tmp", r)

	file := findFile("/tmp/GeoLite2-City*/*.mmdb")
	err = os.Rename(file, "geo.mmdb")
	if err != nil {
		fmt.Println("failed to movefile")
		return
	}
}

func getHash() (string, error) {
	resp, err := http.Get("http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz.md5")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return string(body), err
}

func checkHash() string {
	hash, err := getHash()
	if err != nil {
		fmt.Printf("Error occured getting hash %s", err)
	}
	if hash != models.CurrentHash {
		models.CurrentHash = hash
		fmt.Printf("Saving new hash %s\n", hash)
		getDatabase()
	}
	return hash
}

//GetDatabase gets the database
func GetDatabase() string {
	return checkHash()
}
