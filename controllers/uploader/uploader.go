package uploader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/thedevsaddam/govalidator"
)

// ReadFile consumes *multipart.FileHeader and returns string, error
func ReadFile(file *multipart.FileHeader) (string, error) {

	fileName := file.Filename
	fmt.Println("Reading file: ", fileName)

	fileSize := file.Size
	fmt.Println("File size: ", fileSize)

	data, err := file.Open()
	if err != nil {
		return "There was an error.", err
	}

	src, err := ioutil.ReadAll(data)
	if err != nil {
		return "There was an error.", err
	}
	contents := string(src)

	return contents, nil
}

// CreateTempFile consumes the contents and writes to new file for response
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

// DestroyFile for deletion of tempfile on download request
func DestroyFile(fileLoc string) error {
	err := os.Remove(fileLoc)
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
