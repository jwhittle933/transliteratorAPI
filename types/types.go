package types

import (
	"mime/multipart"
	"os"
)

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
