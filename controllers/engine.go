package Language

import (
	"net/http"

	"github.com/labstack/echo"
)

var mGreek = map[string]string{"α": "a", "β": "b", "γ": "g", "δ": "d", "ε": "e", "ζ": "z", "η": "e", "θ": "th", "ι": "i", "κ": "k", "λ": "l", "μ": "m", "ν": "n", "ξ": "ks", "ο": "o", "π": "p", "ρ": "r", "σ": "s", "ς": "s", "τ": "t", "υ": "y", "φ": "ph", "χ": "ch", "ψ": "ps", "ω": "o"}
var mHebrew = map[string]string{"א": "'", "ב": "b", "ג": "g", "ד": "d", "ה": "h", "ו": "w", "ז": "z", "ח": "ch", "ט": "t", "י": "y", "כ": "k", "ך": "k", "ל": "l", "מ": "m", "ם": "m", "נ": "n", "ן": "n", "ס": "s", "ע": "'", "פ": "p", "ף": "ph", "צ": "ts", "ץ": "ts", "ק": "q", "ר": "r", "שׁ": "s", "שׂ": "sh", "ת": "th"}

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
