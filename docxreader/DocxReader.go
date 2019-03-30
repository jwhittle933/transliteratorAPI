package docxreader

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

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
		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

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
