package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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
	// TODO
	// !! First, copy file to file system
	// !! Then, use os package to read file
	// !! Then, parse contents
	// !! Lastly, write new contents to file and return to client
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "There was an error receiving the file.",
		})
	}
	fmt.Println("Print file: ", file)

	data, err := file.Open()
	if err != nil {
		return err
	}

	src, err := ioutil.ReadAll(data)
	if err != nil {
		return err
	}

	text := string(src)
	fmt.Printf("File data: %s\n", string(text))

	return c.JSON(http.StatusOK, "OK.")
}
