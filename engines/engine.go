package engine

import (
	"net/http"

	"github.com/labstack/echo"
)

type Text struct {
	Body string `json:"body" form:"body" query:"body"`
	Lang string `json:"lang" form:"lang" query:"lang"`
}

func MainHandler(c echo.Context) error {
	lang := c.QueryParam("language")
	text := c.QueryParam("text")
	if len(lang) == 0 || len(text) == 0 {
		return c.String(http.StatusOK, "No language or text submitted.")
	}
	output := transliterate(lang, text)
	return c.JSON(http.StatusOK, output)
}

func transliterate(language string, text string) string {
	var str string
	if language == "Greek" {
		for _, value := range text {
			letter := string(value)
			if letter == " " {
				str += " "
			} else {
				str += mGreek[letter]
			}
		}
	} else if language == "Hebrew" {
		for _, value := range text {
			letter := string(value)
			if letter == " " {
				str += " "
			} else {
				str += mHebrew[letter]
			}
		}
	}
	return str
}

func NewText(c echo.Context) (err error) {
	t := new(Text)
	if err = c.Bind(t); err != nil {
		return
	}
	return c.JSON(http.StatusOK, t)
}
