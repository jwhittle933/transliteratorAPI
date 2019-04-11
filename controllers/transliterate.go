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
	t := new(types.TextSubmission)
	if err := c.Bind(t); err != nil {
		erm = &types.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "Bad submission",
		}
		return c.JSON(http.StatusBadRequest, erm)
	}

	if t.Text == "" {
		erm = &types.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: "No text provided.",
		}
		return c.JSON(http.StatusBadRequest, erm)
	}

	if lang, output := engine.Transliterate(t.Text); output != "Error." {
		response := &types.SuccessfulResponse{
			Code:               http.StatusOK,
			Message:            "Successful.",
			Language:           lang,
			SubmittedText:      t.Text,
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
