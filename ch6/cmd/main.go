package main

import (
	"log"
	"net/http"

	"github.com/holmes89/hello-api/handlers"
	"github.com/holmes89/hello-api/handlers/rest"
	"github.com/holmes89/hello-api/translation"
)

func main() {

	addr := ":8080"

	mux := http.NewServeMux()

	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)
	mux.HandleFunc("/translate/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck) // <1>

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

type Resp struct { // <6>
	Language    string `json:"language"`
	Translation string `json:"translation"`
}
