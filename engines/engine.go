package engine

var mGreek = map[string]string{"α": "a", "Α": "A", "β": "b", "Β": "B", "γ": "g", "Γ": "G", "δ": "d", "Δ": "D", "ε": "e", "Ε": "E", "ζ": "z", "Ζ": "Z", "η": "ē", "Η": "Έ", "θ": "th", "Θ": "Th", "ι": "i", "Ι": "I", "κ": "k", "Κ": "K", "λ": "l", "Λ": "L", "μ": "m", "Μ": "M", "ν": "n", "Ν": "N", "ξ": "ks", "Ξ": "Ks", "ο": "o", "Ο": "O", "π": "p", "Π": "P", "ρ": "r", "Ρ": "R", "ς": "s", "σ": "s", "Σ": "S", "τ": "t", "Τ": "T", "υ": "y", "Υ": "Y", "φ": "ph", "Φ": "Ph", "χ": "ch", "Χ": "Ch", "ψ": "ps", "Ψ": "Ps", "ω": "ō", "Ω": "Ō"}
var greekAccents = map[string][]byte{"᾽": []byte("᾽"), "´": []byte("´"), "`": []byte("`"), "῀": []byte("῀"), "῾": []byte("῾"), "ι": []byte("ι"), "¨": []byte("¨")}
var mHebrew = map[string]string{"א": "'", "ב": "b", "ג": "g", "ד": "d", "ה": "h", "ו": "w", "ז": "z", "ח": "ch", "ט": "t", "י": "y", "כ": "k", "ך": "k", "ל": "l", "מ": "m", "ם": "m", "נ": "n", "ן": "n", "ס": "s", "ע": "'", "פ": "p", "ף": "ph", "צ": "ts", "ץ": "ts", "ק": "q", "ר": "r", "שׁ": "s", "שׂ": "sh", "ת": "th"}

// Transliterate func is the engine of the api.
func Transliterate(text string) (lang string, str string) {
	lang = WhichLang(text)
	if lang := WhichLang(text); lang == "None." {
		return "Unsupported Language.", "Error."
	} else if lang == "Greek" {
		str = ComposeGreekStr(text)
	} else if lang == "Hebrew" {
		str = ComposeHebrewStr(text)
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

// ComposeGreekStr consumes text and transliterates.
func ComposeGreekStr(text string) (str string) {
	for _, value := range text {
		letter := string(value)
		if letter == " " {
			str += letter
		} else {
			str += mGreek[letter]
		}
	}
	return str
}

// ComposeHebrewStr consumes text and transliterates.
func ComposeHebrewStr(text string) (str string) {
	for _, value := range text {
		letter := string(value)
		if letter == " " {
			str += letter
		} else {
			str += mHebrew[letter]
		}
	}
	return str
}
