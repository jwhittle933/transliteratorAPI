package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	engine "../engines"
	"./uploader"
	uuid "github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

// https://www.devdungeon.com/content/working-files-go#everything_is_a_file

// ProcessFile for reading uploaded file
func ProcessFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "There was an error receiving the file.",
		})
	}

	fileContents, bytes, err := uploader.ReadFile(file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "There was an error receiving the file.",
		})
	}

	lang, transliteratedContents := engine.Transliterate(fileContents)
	bytesWritten, pathToFile, err := CreateTempFile(bytes)

	resp := &UploadSuccess{
		Code:               http.StatusOK,
		Message:            "File Succesfully read.",
		Language:           lang,
		OriginalFile:       file,
		TransliteratedText: transliteratedContents,
		BytesWritten:       bytesWritten,
		DownloadLink:       pathToFile,
	}

	return c.JSON(http.StatusOK, resp)
}

// CreateTempFile consumes the contents and writes to new file
// for response
func CreateTempFile(byteSlice []byte) (int, string, error) {
	uuid := uuid.New()
	pathToFile := fmt.Sprintf("./tmp/resp-%d.txt", uuid)

	newFile, err := os.Create(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	openFile, err := os.OpenFile(pathToFile, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer openFile.Close()

	write, err := openFile.Write(byteSlice)

	return write, pathToFile, err
}

// DestroyFile for deletion of tempfile
func DestroyFile(fileLoc string) error {
	err := os.Remove("./tmp/resp.txt")
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// ValidateFile func
func ValidateFile(w http.ResponseWriter, r *http.Request) {
	rules := govalidator.MapData{
		"file:text": []string{"ext:txt, docx, csv", "size:100000", "mime:txt, docx, csv", "required"},
	}

	messages := govalidator.MapData{
		"file:text": []string{"ext:Only txt/docx/csv allowed", "required:document is required"},
	}

	opts := govalidator.Options{
		Request:         r,        // request object
		Rules:           rules,    // rules map
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true,     // all the field to be pass the rules
	}

	v := govalidator.New(opts)
	e := v.Validate()
	err := map[string]interface{}{"validationError": e}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(err)
}
