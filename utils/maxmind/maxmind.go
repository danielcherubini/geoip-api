package maxmind

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/danielcherubini/geoip-api/models"
	"github.com/danielcherubini/geoip-api/services"
	"github.com/danielcherubini/geoip-api/utils"
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

func unTar(filePath string, tempDir string) error {
	fmt.Println("Decompressing Tarball")
	lastIndexOfTempDirString := tempDir[len(tempDir)-1:]
	if lastIndexOfTempDirString != "/" {
		tempDir = tempDir + "/"
	}
	respByte, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("fail to read response data")
		return err
	}
	r := bytes.NewReader(respByte)
	utils.Untar(tempDir, r)
	return nil
}

func copyFileContents(src, dst string) (err error) {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}

func moveDB(tempDir string) error {
	lastIndexOfTempDirString := tempDir[len(tempDir)-1:]
	if lastIndexOfTempDirString != "/" {
		tempDir = tempDir + "/"
	}
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
		return err
	}
	fmt.Println("File download finished")
	return nil
}

func getUrl(urlString string) error {
	var filePath = ""
	urlIndex := strings.Split(urlString, "/")
	filename := urlIndex[len(urlIndex)-1]
	isGzip := strings.Contains(filename, "tar.gz")

	err, filePath := services.DownloadUrl(urlString, filename)
	if err != nil {
		return err
	}

	if isGzip {
		err = unTar(filePath, tempDir)
		if err != nil {
			return err
		}
	}

	err = moveDB(tempDir)
	if err != nil {
		return err
	}
	return nil
}

func getS3(s3Config models.S3Config) error {
	var filePath = ""
	filenameArr := strings.Split(s3Config.Key, "/")
	filename := filenameArr[len(filenameArr)-1]
	isGzip := strings.Contains(filename, "tar.gz")

	err, filePath := services.DownloadS3Url(s3Config, filename)
	if err != nil {
		fmt.Println("error in getS3 from downloadS3Url")
		return err
	}

	if isGzip {
		unTar(filePath, tempDir)
		if err != nil {
			fmt.Println("error in getS3 from unTar")
			return err
		}
	}

	err = moveDB(tempDir)
	if err != nil {
		fmt.Println("error in getS3 from moveDB")
		return err
	}
	return nil
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
func GetDatabase(db models.DBLocation) error {
	fmt.Println("File download starting")
	switch db.Type {
	case "MMDB":
		err := moveDB(db.Location)
		return err
	case "GZDB":
		unTar(db.Location, tempDir)
		err := moveDB(tempDir)
		return err
	case "DBURL":
		err := getUrl(db.Location)
		return err
	case "S3DB":
		err := getS3(db.S3Config)
		return err
	default:
		hash := checkHash()
		checkDBHash(hash)
		return nil
	}
}
