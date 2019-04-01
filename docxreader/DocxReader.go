package docxreader

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// DocxUnzip for reading Word .docx files.
/*
 TODO: Accept []byte of read contents of submitted docx
 TODO: Use read data to create unzip, rather than creating tmp file
 TODO: Write data to disc or keep in memory?
*/
func DocxUnzip(pathToFile, saveLocation string) error {
	reader, err := zip.OpenReader(pathToFile)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(saveLocation, 0755); err != nil {
		return err
	}

	// ExpandDocx(reader) << will be changed

	return nil
}

// ExpandDocx method
// func ExpandDocx(z *zip.ReadCloser) zip.File {

// }

// Unzip for exposing contexts of zip.
// TODO: Modularize
func Unzip(pathToFile, saveLocation string) error {
	reader, err := zip.OpenReader(pathToFile)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(saveLocation, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		/* file.FileInfo().IsDir() returns false in every case
		file.FileInfo() returns os.FileInfo
		*/
		path := filepath.Join(saveLocation, file.Name)
		dirPath := filepath.Dir(path)
		os.MkdirAll(dirPath, 0777)
		_, err := os.Create(path)

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

// XMLManip for manipulating xml
func XMLManip() {
	return
}
