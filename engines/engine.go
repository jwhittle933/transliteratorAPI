package engine

var mGreek = map[string]string{"α": "a", "β": "b", "γ": "g", "δ": "d", "ε": "e", "ζ": "z", "η": "e", "θ": "th", "ι": "i", "κ": "k", "λ": "l", "μ": "m", "ν": "n", "ξ": "ks", "ο": "o", "π": "p", "ρ": "r", "σ": "s", "ς": "s", "τ": "t", "υ": "y", "φ": "ph", "χ": "ch", "ψ": "ps", "ω": "o"}
var mHebrew = map[string]string{"א": "'", "ב": "b", "ג": "g", "ד": "d", "ה": "h", "ו": "w", "ז": "z", "ח": "ch", "ט": "t", "י": "y", "כ": "k", "ך": "k", "ל": "l", "מ": "m", "ם": "m", "נ": "n", "ן": "n", "ס": "s", "ע": "'", "פ": "p", "ף": "ph", "צ": "ts", "ץ": "ts", "ק": "q", "ר": "r", "שׁ": "s", "שׂ": "sh", "ת": "th"}

// Transliterate func is the engine of the api.
func Transliterate(text string) (lang string, str string) {
	lang = OneOf(text)
	if lang == "Greek" {
		// TODO abstract this loop
		for _, value := range text {
			letter := string(value)
			if letter == " " {
				str += " "
			} else {
				str += mGreek[letter]
			}
		}
	} else if lang == "Hebrew" {
		// TODO abstract this loop
		for _, value := range text {
			letter := string(value)
			if letter == " " {
				str += " "
			} else {
				str += mHebrew[letter]
			}
		}
	}

	return lang, str
}

// LauguageCheck for determining language and text correspond.
func LauguageCheck(lang string, text string) bool {
	return true
}

// OneOf used for language autodetect.
func OneOf(text string) string {
	for _, val := range text {
		if x := mGreek[string(val)]; x != "" {
			return "Greek"
		}
	}
	return "Hebrew"
}
