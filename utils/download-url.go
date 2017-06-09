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
		fmt.Println("Missing slash at end")
		tempDir = tempDir + "/"
	}

	filePath = tempDir + filename

	fmt.Println(urlString)
	fmt.Println(filePath)

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

	n, err := io.Copy(output, response.Body)
	if err != nil {
		fmt.Println("Error while downloading", urlString, "-", err)
		return err, ""
	}

	fmt.Println(n, "bytes downloaded.")
	return nil, filePath
}
