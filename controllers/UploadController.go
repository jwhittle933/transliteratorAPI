package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	engine "../engines"
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
	if err := CreateTempFile(src); err != nil {
		errCheck(c, err)
	}

	// lang of type string, transliteratedContents of type string
	lang, transliteratedContents := engine.Transliterate(string(src))

	errCheck(c, err)

	return c.JSON(http.StatusOK, &UploadSuccess{
		Code:               http.StatusOK,
		Message:            "File Succesfully read.",
		Language:           lang,
		OriginalFile:       data,
		TempFile:           tempFile,
		FileType:           mime,
		TransliteratedText: transliteratedContents,
		BytesWritten:       len(src),
		// DownloadLink:       "http://localhost:3000" + tempFile,
	})
}

// CreateTempFile wrapper func for ioutil.TempFile
func CreateTempFile(src []byte) error {
	uuid := uuid.New()
	tempFile, err := ioutil.TempFile("../tmp", fmt.Sprintf("%d", uuid))
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(src); err != nil {
		log.Fatal(err)
		return err
	}

	if err := tempFile.Close(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
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
