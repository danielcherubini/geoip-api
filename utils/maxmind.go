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
	fmt.Println("File download finished")
}

func getUrl(urlString string) {
	var filePath = ""
	urlIndex := strings.Split(urlString, "/")
	filename := urlIndex[len(urlIndex)-1]
	isGzip := strings.Contains(filename, "tar.gz")

	err, filePath := downloadUrl(urlString, filename)
	if err != nil {
		return
	}

	if isGzip {
		unTar(filePath, tempDir)
	}

	moveDB(tempDir)
}

func getS3(s3Config models.S3Config) {
	var filePath = ""
	filenameArr := strings.Split(s3Config.Key, "/")
	filename := filenameArr[len(filenameArr)-1]
	isGzip := strings.Contains(filename, "tar.gz")

	err, filePath := downloadS3Url(s3Config, filename)
	if err != nil {
		return
	}

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
	fmt.Println("File download starting")
	switch db.Type {
	case "MMDB":
		moveDB(db.Location)
		break
	case "GZDB":
		unTar(db.Location, tempDir)
		moveDB(tempDir)
		break
	case "DBURL":
		getUrl(db.Location)
		break
	case "S3DB":
		getS3(db.S3Config)
		break
	default:
		hash := checkHash()
		checkDBHash(hash)
		break
	}
}
