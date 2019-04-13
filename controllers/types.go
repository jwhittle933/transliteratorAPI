package controllers

import (
	"mime/multipart"
	"os"
)

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

// TextSubmission struct
type TextSubmission struct {
	Text string `json:"text" form:"text" query:"text"`
}

// User struct for registering users
type User struct {
	ID        int    `json:"id" form:"id" query:"id"`
	FirstName string `json:"firstname" form:"firstname" query:"firstname"`
	LastName  string `json:"lastname" form:"lastname" query:"lastname"`
	Email     string `json:"email" form:"email" query:"email"`
	Pass      string `json:"password" form:"password" query:"password"` //encrypt
}
