package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/holmes89/hello-api/translation"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	language := r.URL.Query().Get("language")
	if language == "" {
		language = "english"
	}
	fmt.Println(r.URL.Path)
	word := strings.ReplaceAll(r.URL.Path, "/translate/", "")
	fmt.Println(word)
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
