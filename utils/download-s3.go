package utils

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/danmademe/geoip-api/models"
)

func downloadS3Url(s3Config models.S3Config, filename string) (err error, filePath string) {
	filePath = tempDir + filename
	file, err := os.Create(filePath)

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
