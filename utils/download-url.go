package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadUrl(urlString string, filename string) (err error, filePath string) {

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
