package docxreader

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// Zip struct for handling extention methods on unziped files.
type Zip struct {
	Reader *zip.ReadCloser
	Files  []*zip.File
}

// XMLData struct for Unmarshalling xml.
type XMLData struct {
	Text string `xml:"w:t"`
}

// Document type for storing word/document.xml extracted from docx zip.
type Document struct {
	Doc *zip.File
}

// DocxUnzip for reading Word .docx files.
/*
 TODO: Accept []byte of read contents of submitted docx
 TODO: Use read data to create unzip, rather than creating tmp file
 TODO: Write data to disc or keep in memory?
*/
func DocxUnzip(pathToFile string) error {
	saveLocation := fmt.Sprintf("./saves/dir-%d/unzip", uuid.New())
	zip := ExtractFiles(pathToFile)

	if err := os.MkdirAll(saveLocation, 0755); err != nil {
		return err
	}

	if err := zip.MapFiles(saveLocation); err != nil {
		return err
	}

	return nil
}

// ExtractFiles returns []*zip.File.
func ExtractFiles(pathToFile string) *Zip {
	reader, err := zip.OpenReader(pathToFile)
	if err != nil {
		panic(err)
	}
	// defer reader.Close()
	return &Zip{
		Reader: reader,
		Files:  reader.File,
	}
}

// MapFiles for iterating through zip.File slice
// and performing an operation on it.
func (f *Zip) MapFiles(saveLocation string) error {
	for _, file := range f.Files {
		path := filepath.Join(saveLocation, file.Name)
		dirPath := filepath.Dir(path)
		os.MkdirAll(dirPath, 0777)
		_, err := os.Create(path)
		if err != nil {
			return err
		}

		if err := CopyToOS(file, path); err != nil {
			return err
		}
	}
	return nil
}

// CopyToOS for mapping over files.
func CopyToOS(file *zip.File, filePath string) error {
	fileReader, err := file.Open()
	if err != nil {
		return err
	}
	defer fileReader.Close()

	targetFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}

	defer targetFile.Close()

	if _, err := io.Copy(targetFile, fileReader); err != nil {
		return err
	}

	// t := XMLData{}
	// data, err := ioutil.ReadAll(targetFile)
	// if err != nil {
	// 	return err
	// }
	// xml.Unmarshal([]byte(data), &t)
	// fmt.Println("Reading XML", t.Text)

	return nil
}

//FindDoc locates file named word/document.xml
func (f *Zip) FindDoc() (file *Document) {
	for _, fi := range f.Files {
		if fi.Name == "word/document.xml" {
			file = &Document{
				Doc: fi,
			}
		}
	}
	return file
}

// XMLExtractText for manipulating xml
func XMLExtractText() {
	return
}
