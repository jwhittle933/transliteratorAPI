package engine

var mGreek = map[string]string{"α": "a", "β": "b", "γ": "g", "δ": "d", "ε": "e", "ζ": "z", "η": "e", "θ": "th", "ι": "i", "κ": "k", "λ": "l", "μ": "m", "ν": "n", "ξ": "ks", "ο": "o", "π": "p", "ρ": "r", "σ": "s", "ς": "s", "τ": "t", "υ": "y", "φ": "ph", "χ": "ch", "ψ": "ps", "ω": "o"}
var mHebrew = map[string]string{"א": "'", "ב": "b", "ג": "g", "ד": "d", "ה": "h", "ו": "w", "ז": "z", "ח": "ch", "ט": "t", "י": "y", "כ": "k", "ך": "k", "ל": "l", "מ": "m", "ם": "m", "נ": "n", "ן": "n", "ס": "s", "ע": "'", "פ": "p", "ף": "ph", "צ": "ts", "ץ": "ts", "ק": "q", "ר": "r", "שׁ": "s", "שׂ": "sh", "ת": "th"}

// Transliterate func is the engine of the api.
func Transliterate(text string) (lang string, str string) {
	// TODO
	// !! First, copy file to file system
	// !! Then, use os package to read file
	// !! Then, parse contents
	// !! Lastly, write new contents to file and return to client
	lang = WhichLang(text)
	if lang == "None." {
		return "Unsupported Language.", "Error."
	}
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

// WhichLang used for language autodetect.
func WhichLang(text string) string {
	for _, val := range text {
		if x := mGreek[string(val)]; x != "" {
			return "Greek"
		}
	}
	for _, val := range text {
		if x := mHebrew[string(val)]; x != "" {
			return "Hebrew"
		}
	}
	return "None."
}
