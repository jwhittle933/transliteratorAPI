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

// Transliterator route handler
func Transliterator(c echo.Context) error {
	var erm *ErrorMessage
	lang := c.QueryParam("language")
	text := c.QueryParam("text")
	if len(lang) == 0 {
		erm = &ErrorMessage{
			Code:    400,
			Message: "No language specified.",
		}
		return c.JSON(http.StatusBadRequest, erm)
	}
	if len(text) == 0 {
		erm = &ErrorMessage{
			Code:    400,
			Message: "No text provided.",
		}
		return c.JSON(http.StatusBadRequest, erm)
	}
	output := engine.Transliterate(lang, text)
	return c.JSON(http.StatusOK, output)
}
