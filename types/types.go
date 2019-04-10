package types

import (
	"mime/multipart"
	"os"
)

// User type for data persistenceß
type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

// UserFirst struct for registering users
type UserFirst struct {
	FirstName string `json:"firstname" form:"firstname" query:"firstname"`
	LastName  string `json:"lastname" form:"lastname" query:"lastname"`
	Email     string `json:"email" form:"email" query:"email"`
	Pass      string `json:"password" form:"password" query:"password"` //encrypt
}

// Resp struct for basic response schema.
type Resp struct {
	Code    int64
	Message string
}

// ErrorMessage for error repsonses
type ErrorMessage struct {
	Code    int64
	Message string
}

// SuccessfulResponse for transliterate success message.
type SuccessfulResponse struct {
	Code               int64
	Message            string
	Language           string
	SubmittedText      string
	TransliteratedText string
}

// UploadSuccess for upload success message.
type UploadSuccess struct {
	Code               int64
	Message            string
	Language           string
	OriginalFile       multipart.File
	TempFile           *os.File
	FileType           string
	TransliteratedText string
	BytesWritten       int
	DownloadLink       string
}

// File struct
type File struct {
	File        multipart.File
	Mime        string
	Size        int64
	ByteContent []byte
	TextContext string
}
