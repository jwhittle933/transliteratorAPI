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
