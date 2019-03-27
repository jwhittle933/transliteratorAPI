package uploader

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
)

// ReadFile consumes *multipart.FileHeader and returns string, error
func ReadFile(file *multipart.FileHeader) (string, error) {

	fileName := file.Filename
	fmt.Println("Reading file: ", fileName)

	fileSize := file.Size
	fmt.Println("File size: ", fileSize)

	data, err := file.Open()
	if err != nil {
		return "There was an error.", err
	}

	src, err := ioutil.ReadAll(data)
	if err != nil {
		return "There was an error.", err
	}
	contents := string(src)

	return contents, nil
}
