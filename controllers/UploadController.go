package controllers

import (
	"encoding/json"
	"log"
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

// ProcessFile for reading uploaded file
func ProcessFile(c echo.Context) error {
	var (
		newFile *os.File
		err     error
	)

	newFile, err = os.Create("test.txt")
	check(err)
	log.Println(newFile)
	newFile.Close()

	// data, err := ioutil.ReadFile("test.txt")
	// check(err)
	// fmt.Println(data)

	// f, err := os.Open("test.txt")
	// check(err)

	// b1 := make([]byte, 5)
	// n1, err := f.Read(b1)
	// check(err)
	// fmt.Printf("%d bytes: %s\n", n1, string(b1))

	// read := bufio.NewReader(f)
	// b4, err := read.Peek(5)
	// check(err)
	// fmt.Printf("5 bytes: %s\n", string(b4))

	return c.JSON(http.StatusOK, "OK.")
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
