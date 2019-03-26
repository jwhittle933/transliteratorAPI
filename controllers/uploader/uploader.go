package uploader

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
)

func handleFile(file *multipart.FileHeader) (string, error) {

	fileName := file.Filename
	fmt.Println("Reading file: ", fileName)

	fileSize := file.Size
	fmt.Println("File size: ", fileSize)

	data, err := file.Open()
	if err != nil {
		return "", err
	}

	src, err := ioutil.ReadAll(data)
	if err != nil {
		return "", err
	}
	contents := string(src)
	fmt.Printf("File data: %s\n: ", src)

	return contents, err
}
