package pdfreader

import (
	"fmt"
	"mime/multipart"
)

//PdfReader main func for reading pdf
func PdfReader(f *multipart.FileHeader) ([]byte, error) {
	file, err := f.Open()

	if err != nil {
		return []byte{}, err
	}

	defer file.Close()

	buffer := make([]byte, 9999)
	file.Read(buffer)
	fmt.Println(string(buffer))

	return buffer, nil
}
