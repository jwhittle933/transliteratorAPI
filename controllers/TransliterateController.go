package controllers

import (
	"net/http"

	engine "../engines"
	"github.com/labstack/echo"
)

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
	if lang, output := engine.Transliterate(text); output != "Error." {
		response := &SuccessfulResponse{
			Code:               http.StatusOK,
			Message:            "Successful.",
			Language:           lang,
			SubmittedText:      text,
			TransliteratedText: output,
		}
		return c.JSON(http.StatusOK, response)
	}
	response := &ErrorMessage{
		Code:    http.StatusBadRequest,
		Message: "Error",
	}
	return c.JSON(http.StatusBadRequest, response)
}
