package engine

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
		if x := Greek[string(val)]; x != "" {
			return "Greek"
		}
	}
	for _, val := range text {
		if x := Hebrew[string(val)]; x != "" {
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
			str += Greek[letter]
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
			str += Hebrew[letter]
		}
	}
	return str
}
