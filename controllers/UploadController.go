package controllers

import (
	"encoding/json"
	"net/http"

	engine "../engines"
	"./uploader"
	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

// https://www.devdungeon.com/content/working-files-go#everything_is_a_file

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

	fileContents, err := uploader.ReadFile(file)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "There was an error receiving the file.",
		})
	}

	lang, transliteratedContents := engine.Transliterate(fileContents)

	resp := &SuccessfulResponse{
		Code:               http.StatusOK,
		Message:            "File Succesfully read.",
		Language:           lang,
		SubmittedText:      fileContents,
		TransliteratedText: transliteratedContents,
	}

	return c.JSON(http.StatusOK, resp)
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
