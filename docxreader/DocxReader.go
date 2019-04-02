package docxreader

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// ZipFiles struct for handling extention methods on unziped files.
type ZipFiles struct {
	Files []*zip.File
}

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

	if err := ExtractFiles(reader).MapFiles(saveLocation); err != nil {
		return err
	}

	return nil
}

// ExtractFiles returns []*zip.File.
func ExtractFiles(z *zip.ReadCloser) *ZipFiles {
	return &ZipFiles{
		Files: z.File,
	}
}

// MapFiles for iterating through zip.File slice
// and performing an operation on it.
func (f *ZipFiles) MapFiles(saveLocation string) error {
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

	return nil
}

// XMLManip for manipulating xml
func XMLManip() {
	return
}
