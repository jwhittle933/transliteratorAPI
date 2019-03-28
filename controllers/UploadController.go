package controllers

import (
	"net/http"

	engine "../engines"
	"./uploader"
	"github.com/labstack/echo"
)

// https://www.devdungeon.com/content/working-files-go#everything_is_a_file

// UploadHandler experimental http handler
func UploadHandler(w http.ResponseWriter, r *http.Request) {

}

// Uploader for reading uploaded file
func Uploader(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "There was an error receiving the file.",
		})
	}

	// openedFile, err := file.Open()

	fileContents, err := uploader.ReadFile(file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "There was an error receiving the file.",
		})
	}

	lang, transliteratedContents := engine.Transliterate(fileContents)
	f, bytesWritten, pathToFile, err := uploader.CreateTempFile([]byte(transliteratedContents))

	resp := &UploadSuccess{
		Code:               http.StatusOK,
		Message:            "File Succesfully read.",
		Language:           lang,
		OriginalFile:       f,
		TransliteratedText: transliteratedContents,
		BytesWritten:       bytesWritten,
		DownloadLink:       "http://localhost:3000" + pathToFile,
	}

	return c.JSON(http.StatusOK, resp)
}
