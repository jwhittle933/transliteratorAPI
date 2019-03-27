package controllers

import "os"

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
	Code         int64
	Message      string
	Language     string
	OriginalFile *os.File
	DownloadLink *os.File
}
