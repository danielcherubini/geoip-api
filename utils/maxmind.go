package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/danmademe/geoip-api/models"
)

var DBUrl string
var tempDir = os.TempDir()

func findFile(pattern string) string {

	matches, err := filepath.Glob(pattern)

	if err != nil {
		fmt.Println(err)
	}

	if len(matches) < 1 {
		return ""
	}
	return matches[0]
}

func unTar(filePath string, tempDir string) {
	fmt.Println("Decompressing Tarball")
	respByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("fail to read response data")
		return
	}
	r := bytes.NewReader(respByte)
	Untar(tempDir, r)
}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer func() {
		cerr := out.Close()
		if err == nil {
			err = cerr
		}
	}()
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}

func moveDB(tempDir string) {
	var file string
	if strings.Contains(tempDir, ".mmdb") {
		file = tempDir
	} else {
		file = findFile(tempDir + "*.mmdb")
		if file == "" {
			file = findFile(tempDir + "*/*.mmdb")
		}
	}
	err := copyFileContents(file, "geo.mmdb")
	if err != nil {
		fmt.Println("failed to movefile")
		os.Exit(1)
	}
}

func getDatabase(url string) {
	urlIndex := strings.Split(url, "/")
	filename := urlIndex[len(urlIndex)-1]
	isGzip := strings.Contains(filename, "tar.gz")

	filePath := tempDir + filename
	fmt.Println(url)
	fmt.Println(filePath)

	output, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error while creating", filePath, "-", err)
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

	if isGzip {
		unTar(filePath, tempDir)
	}
	moveDB(tempDir)
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

func checkDBHash(hash string) {
	if hash != models.CurrentHash {
		models.CurrentHash = hash
		fmt.Printf("Saving new hash %s\n", hash)
	}
}

func checkHash() string {
	hash, err := getHash()
	if err != nil {
		fmt.Printf("Error occured getting hash %s", err)
	}
	return hash
}

//GetDatabase gets the database
func GetDatabase(db models.DBLocation) {
	switch db.Type {
	case "MMDB":
		moveDB(db.Location)
		break
	case "GZDB":
		unTar(db.Location, tempDir)
		moveDB(tempDir)
		break
	case "DBURL":
		getDatabase(db.Location)
		break
	default:
		hash := checkHash()
		checkDBHash(hash)
		break
	}
}
