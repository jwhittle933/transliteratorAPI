package controllers

import (
	"net/http"

	engine "../engines"
	"github.com/labstack/echo"
)

type Error struct {
	code    int64
	message string
}

type Success struct {
	language           string
	submittedText      string
	transliteratedText string
}

func Transliterator(c echo.Context) error {
	lang := c.QueryParam("language")
	text := c.QueryParam("text")
	if len(lang) == 0 {
		erm := &Error{
			code:    400,
			message: "No language specified.",
		}
		return c.JSON(http.StatusOK, erm)
	}
	if len(text) == 0 {
		erm := &Error{
			code:    400,
			message: "No text specified.",
		}
		return c.JSON(http.StatusOK, erm)
	}
	output := engine.Transliterate(lang, text)
	return c.JSON(http.StatusOK, output)
}
