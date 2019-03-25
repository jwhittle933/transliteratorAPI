package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

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

// https://www.devdungeon.com/content/working-files-go#everything_is_a_file
// https://golang.org/pkg/os/#File.Close

// ProcessFile for reading uploaded file
func ProcessFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "There was an error receiving the file.",
		})
	}
	fmt.Println("Print file: ", file)

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "There was an error opening your file.",
		})
	}
	fmt.Println("Print src: ", src)

	dst, err := os.Create(file.Filename)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "There was an error accessing your file.",
		})
	}
	fmt.Println("Print dst: ", dst)

	readFile, err := os.Open(file.Filename)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "There was an error reading your file.",
		})
	}
	b := make([]byte, 10)
	f, err := readFile.Read(b)
	fmt.Println("Print f: ", f)
	fileInfo, err := readFile.Stat()
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	bytespread, err := readFile.Read(buffer)
	fmt.Println("FileInfo: ", fileInfo)
	fmt.Println("FileSize: ", fileSize)
	fmt.Println("Bytes: ", bytespread)
	fmt.Println("Buffer: ", buffer)

	defer src.Close()
	defer readFile.Close()
	defer dst.Close()

	return c.JSON(http.StatusOK, "OK.")
}
