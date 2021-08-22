package translation

import "strings"

func Translate(word string, language string) string {
	word = sanitizeInput(word)         // <1>
	language = sanitizeInput(language) // <2>

	if word != "hello" {
		return ""
	}

	switch language {
	case "english":
		return "hello"
	case "finnish":
		return "hei"
	case "german":
		return "hallo"
	default:
		return ""
	}
}

func sanitizeInput(w string) string { // <3>
	w = strings.ToLower(w)
	return strings.TrimSpace(w)
}
