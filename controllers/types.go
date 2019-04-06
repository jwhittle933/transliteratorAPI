package controllers

import (
	"mime/multipart"
	"os"
)

// Resp struct for response schema.
type Resp struct {
	Code    int64
	Message string
}

// ErrorMessage for forming error repsonses
type ErrorMessage struct {
	Code    int64
	Message string
}

// SuccessfulResponse struct.
type SuccessfulResponse struct {
	Code               int64
	Message            string
	Language           string
	SubmittedText      string
	TransliteratedText string
}

// UploadSuccess struct.
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
