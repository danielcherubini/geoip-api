package services

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/danielcherubini/geoip-api/models"
)

var tempDir = os.TempDir()

//DownloadS3Url takes a S3Config object and a filename string and gets the file from S3
func DownloadS3Url(s3Config models.S3Config, filename string) (err error, filePath string) {
	lastIndexOfTempDirString := tempDir[len(tempDir)-1:]
	if lastIndexOfTempDirString != "/" {
		tempDir = tempDir + "/"
	}

	filePath = tempDir + filename
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error making filepath")
		return err, ""
	}
	// The session the S3 Downloader will use
	sess := session.Must(session.NewSession(&aws.Config{Region: aws.String(s3Config.Region)}))

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)
	input := &s3.GetObjectInput{
		Bucket: aws.String(s3Config.Bucket),
		Key:    aws.String(s3Config.Key),
	}

	_, err = downloader.Download(file, input)
	if err != nil {
		fmt.Println("S3 Downloader Error")
		return err, ""
	}
	defer file.Close()

	return nil, file.Name()
}

//DownloadUrl takes a url string an filename string and gets the file from the web
func DownloadUrl(urlString string, filename string) (err error, filePath string) {

	lastIndexOfTempDirString := tempDir[len(tempDir)-1:]
	if lastIndexOfTempDirString != "/" {
		tempDir = tempDir + "/"
	}

	filePath = tempDir + filename

	output, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error while creating", filePath, "-", err)
		return err, ""
	}
	defer output.Close()

	response, err := http.Get(urlString)
	if err != nil {
		fmt.Println("Error while downloading", urlString, "-", err)
		return err, ""
	}
	defer response.Body.Close()

	_, err = io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", urlString, "-", err)
		return err, ""
	}

	return nil, filePath
}
