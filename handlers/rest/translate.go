// Package rest houses all rest handlers for the application.
package rest

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/holmes89/hello-api/translation"
)

// Resp is the response sent back to the user.
type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

// TranslateHandler will take a given request with a path value of the
// word to be translated and a query parameter of the language to translate to.
func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = "english"
	}

	word := filepath.Base(r.URL.Path)
	translation := translation.Translate(word, language)
	if translation == "" {
		language = ""
		w.WriteHeader(404)
	}
	resp := Resp{
		Language:    language,
		Translation: translation,
	}
	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
