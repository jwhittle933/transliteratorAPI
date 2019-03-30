package docxreader

import (
	"fmt"

	"github.com/nguyenthenguyen/docx"
)

// NOTES :
// Office Docs (docx, xlsx, *x) are just zip files with xml.

// DocxReader for reading Word .docx files.
func DocxReader() {
	r, _ := docx.ReadDocxFile("../testfiles/hebrew.docx")

	docx := r.Editable()
	fmt.Println(docx)

	defer r.Close()

	// buffer := make([]byte, 9999)

	// file, err := f.Open()
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// int, err := file.ReadAt(buffer, 45)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(int)
}
