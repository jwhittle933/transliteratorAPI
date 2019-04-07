package controllers

import (
	"net/http"

	"github.com/jwhittle933/transliteratorAPI/engine"
	"github.com/jwhittle933/transliteratorAPI/types"
	"github.com/labstack/echo"
)

// TransliterateController route handler
func TransliterateController(c echo.Context) error {
	var erm *types.ErrorMessage
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
