package controllers

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	engine "../engines"
	"./uploader"
	"github.com/google/uuid"
	"github.com/labstack/echo"
)

// https://www.devdungeon.com/content/working-files-go#everything_is_a_file

// Uploader for reading uploaded file
func Uploader(c echo.Context) error {
	// file of type multitpart.FileHeader
	file, err := c.FormFile("file")
	errCheck(c, err)

	data, err := file.Open()
	errCheck(c, err)
	defer data.Close()

	src, err := ioutil.ReadAll(data)
	errCheck(c, err)

	mime := http.DetectContentType(src)

	uuid := uuid.New()
	tempFile, err := ioutil.TempFile("../tmp", fmt.Sprintf("%d", uuid))
	defer os.Remove(tempFile.Name())

	// f of type os.File, bytesWritten of type int, pathToFile of type string
	// TODO: replace with ioutil.TempFile <<<<
	f, bytesWritten, pathToFile, err := uploader.CreateTempFile([]byte(string(src)), "txt")

	// lang of type string, transliteratedContents of type string
	lang, transliteratedContents := engine.Transliterate(string(src))

	errCheck(c, err)

	return c.JSON(http.StatusOK, &UploadSuccess{
		Code:               http.StatusOK,
		Message:            "File Succesfully read.",
		Language:           lang,
		OriginalFile:       f,
		FileType:           mime,
		TransliteratedText: transliteratedContents,
		BytesWritten:       bytesWritten,
		DownloadLink:       "http://localhost:3000" + pathToFile,
	})
}

// DetectFileType to determine mime
func DetectFileType(file multipart.File) (string, error) {
	src, err := ioutil.ReadAll(file)
	if err != nil {
		return "There was an error reading the file.", err
	}
	mime := http.DetectContentType(src)
	return mime, nil
}

func errCheck(c echo.Context, err error) error {
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "There was an error.",
		})
	}
	return nil
}
