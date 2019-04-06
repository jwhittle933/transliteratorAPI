package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/jwhittle933/docxology"
	"github.com/labstack/echo"
)

// https://www.devdungeon.com/content/working-files-go#everything_is_a_file

// UploadController for reading uploaded file
func UploadController(c echo.Context) error {
	var transliteratedContents string
	var lang string

	file, err := c.FormFile("file")
	errCheck(c, err)

	data, err := file.Open()
	errCheck(c, err)
	defer data.Close()

	src, err := ioutil.ReadAll(data)
	errCheck(c, err)

	mime := http.DetectContentType(src)

	if mime == "application/xml" {
		zip := docxology.ExtractFileHTTP(file)
		zipFile := zip.FindDoc("word/document.xml")
		macroData := zipFile.XMLExtractText()
		documentText := macroData.Text
		lang, transliteratedContents = Transliterate(documentText)
	}

	if mime == "application/pdf" {
		// Extract the pdf text
	}

	if mime == "text/plain" {
		lang, transliteratedContents = Transliterate(string(src))
	}

	tempFile, err := CreateTempFile(src)
	if err != nil {
		errCheck(c, err)
	}

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

// TransliterateController route handler
func TransliterateController(c echo.Context) error {
	var erm *ErrorMessage
	req := c.Request()
	fmt.Println(req)
	text := c.QueryParam("text")
	if len(text) == 0 {
		erm = &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "No text provided.",
		}
		return c.JSON(http.StatusBadRequest, erm)
	}
	if lang, output := Transliterate(text); output != "Error." {
		response := &SuccessfulResponse{
			Code:               http.StatusOK,
			Message:            "Successful.",
			Language:           lang,
			SubmittedText:      text,
			TransliteratedText: output,
		}
		return c.JSON(http.StatusOK, response)
	}
	erm = &ErrorMessage{
		Code:    http.StatusBadRequest,
		Message: "Error",
	}
	return c.JSON(http.StatusBadRequest, erm)
}

// CreateTempFile wrapper func for ioutil.TempFile
func CreateTempFile(src []byte) (*os.File, error) {
	uuid := uuid.New()
	tempDir := "../tmp"
	tempFileName := "file-" + fmt.Sprintf("%d", uuid)
	// TODO: create download path-to-file
	tempFile, err := ioutil.TempFile(tempDir, tempFileName)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer os.Remove(tempFile.Name())

	if _, err := tempFile.Write(src); err != nil {
		log.Fatal(err)
		return nil, err
	}

	if err := tempFile.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return tempFile, nil
}

// DestroyFile for deletion of tempfile on download request
func DestroyFile(fileLoc string) error {
	err := os.Remove(fileLoc)
	if err != nil {
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