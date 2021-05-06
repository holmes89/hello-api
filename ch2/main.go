package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	addr := ":8080"
	defaultLanguage := "English"

	mux := http.NewServeMux()

	log.Printf("using default language: %s", defaultLanguage)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		resp := Resp{
			Language:    defaultLanguage,
			Translation: langs[defaultLanguage],
		}
		if err := enc.Encode(resp); err != nil {
			panic("unable to encode response")
		}
	})

	fmt.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

var langs = map[string]string{
	"English": "Hello",
	"Finnish": "Hei",
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
