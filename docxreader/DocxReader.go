package docxreader

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Body struct
type Body struct {
	Paragraph []string `xml:"p>r>t"`
}

// Document struct
type Document struct {
	XMLName xml.Name `xml:"document"`
	Body    Body     `xml:"body"`
}

// NOTES :
// Office Docs (docx, xlsx, *x) are just zip files with xml.

// DocxReader for reading Word .docx files.
func DocxReader() {
	//
}

// Unzip for exposing contexts of zip.
func Unzip(pathToFile, saveLocation string) error {
	reader, err := zip.OpenReader(pathToFile)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(saveLocation, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(saveLocation, file.Name)

		/* file.FileInfo().IsDir() returns false in every case
		file.FileInfo() returns os.FileInfo
		*/
		fmt.Println("FileInfo(): ", file.FileInfo())

		dirPath := filepath.Dir(path)
		os.MkdirAll(dirPath, 0777)
		_, err := os.Create(path)

		fmt.Println("Opening ", file)
		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
