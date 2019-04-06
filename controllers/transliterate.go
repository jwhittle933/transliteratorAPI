package controllers

import (
	"fmt"
	"net/http"

	engine "../engine"
	"../types"
	"github.com/labstack/echo"
)

// Transliterator route handler
func Transliterator(c echo.Context) error {
	var erm *types.ErrorMessage
	req := c.Request()
	fmt.Println(req)
	text := c.QueryParam("text")
	if len(text) == 0 {
		erm = &types.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "No text provided.",
		}
		return c.JSON(http.StatusBadRequest, erm)
	}
	if lang, output := engine.Transliterate(text); output != "Error." {
		response := &types.SuccessfulResponse{
			Code:               http.StatusOK,
			Message:            "Successful.",
			Language:           lang,
			SubmittedText:      text,
			TransliteratedText: output,
		}
		return c.JSON(http.StatusOK, response)
	}
	erm = &types.ErrorMessage{
		Code:    http.StatusBadRequest,
		Message: "Error",
	}
	return c.JSON(http.StatusBadRequest, erm)
}
