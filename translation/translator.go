// Package translation provides the ability to translate words between languages.
package translation

import "strings"

// Translate a given word to a the passed in language.
func Translate(word string, language string) string {
	word = sanitizeInput(word)
	language = sanitizeInput(language)

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
	case "french":
		return "bonjour"
	default:
		return ""
	}

}

func sanitizeInput(w string) string {
	w = strings.ToLower(w)
	return strings.TrimSpace(w)
}
