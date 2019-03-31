package uploader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"../../docxreader"
	"../../pdfreader"
	"github.com/google/uuid"
	"github.com/thedevsaddam/govalidator"
)

// ReadFile consumes *multipart.FileHeader and returns string, error
func ReadFile(file *multipart.FileHeader) (string, string, error) {
	fmt.Println("Reading file (from ReadFile): ", file.Filename)
	fmt.Println("File size (from ReadFile): ", file.Size)

	data, err := file.Open()
	if err != nil {
		return "", "There was an error.", err
	}

	src, err := ioutil.ReadAll(data)
	if err != nil {
		return "", "There was an error.", err
	}

	mimeType := http.DetectContentType(src)
	contents := string(src)

	if mimeType == "application/pdf" {
		// pdfreader.PdfReader from package pdfreader >> Experimental
		pdfFileBytes, _ := pdfreader.PdfReader(file)
		fmt.Println("PDF Detected. BYTE READER: ", pdfFileBytes)
	}

	if mimeType == "application/zip" {
		_, _, pathToFile, err := CreateTempFile(src, "docx")
		if err != nil {
			panic(err)
		}
		if err := docxreader.Unzip(pathToFile, "./testfiles/unzip"); err != nil {
			fmt.Println(err)
		}
	}

	return mimeType, contents, nil
}

// CreateTempFile consumes the contents and writes to new file for response
func CreateTempFile(byteSlice []byte, mime string) (*os.File, int, string, error) {
	uuid := uuid.New()
	pathToFile := fmt.Sprintf("./tmp/file-%d.%s", uuid, mime)

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

	return newFile, write, pathToFile, err
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
		"file:text": []string{"ext:txt, docx, pdf", "size:100000", "mime:txt, docx, pdf", "required"},
	}

	messages := govalidator.MapData{
		"file:text": []string{"ext:Only txt/docx/pdf allowed", "required:document is required"},
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
