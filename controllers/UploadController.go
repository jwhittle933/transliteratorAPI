package controllers

import (
	"net/http"

	engine "../engines"
	pdfreader "../pdfreader"
	"./uploader"
	"github.com/labstack/echo"
)

// https://www.devdungeon.com/content/working-files-go#everything_is_a_file

// Uploader for reading uploaded file
func Uploader(c echo.Context) error {
	// file of type multitpart.FileHeader
	file, err := c.FormFile("file")
	errCheck(c, err)

	multiFile, err := pdfreader.PdfReader(file)

	// mime of type string, fileContents of type string
	mime, fileContents, err := uploader.ReadFile(file)
	errCheck(c, err)

	// lang of type string, transliteratedContents of type string
	lang, transliteratedContents := engine.Transliterate(fileContents)

	// f of type os.File, bytesWritten of type int, pathToFile of type string
	f, bytesWritten, pathToFile, err := uploader.CreateTempFile([]byte(transliteratedContents))
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

func errCheck(c echo.Context, err error) error {
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "There was an error.",
		})
	}
	return nil
}
