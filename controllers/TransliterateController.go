package controllers

import (
	"net/http"

	engine "../engines"
	"github.com/labstack/echo"
)

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

// Transliterator route handler
func Transliterator(c echo.Context) error {
	var erm *ErrorMessage
	text := c.QueryParam("text")
	if len(text) == 0 {
		erm = &ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "No text provided.",
		}
		return c.JSON(http.StatusBadRequest, erm)
	}
	lang, output := engine.Transliterate(text)
	response := &SuccessfulResponse{
		Code:               http.StatusOK,
		Message:            "Successful.",
		Language:           lang,
		SubmittedText:      text,
		TransliteratedText: output,
	}
	return c.JSON(http.StatusOK, response)
}
