package engine

import (
	"net/http"

	"github.com/labstack/echo"
)

func MainHandler(c echo.Context) error {
	lang := c.QueryParam("language")
	text := c.QueryParam("text")
	if len(lang) == 0 || len(text) == 0 {
		return c.String(http.StatusOK, "No language or text submitted.")
	}
	output := transliterate(lang, text)
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{
		"output": output,
	})
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
