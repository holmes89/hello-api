package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/holmes89/hello-api/handlers"
	"github.com/holmes89/hello-api/handlers/rest"
	"github.com/holmes89/hello-api/translation"
)

func main() {
	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if addr == ":" {
		addr = ":8080" // <1>
	}

	mux := http.NewServeMux()

	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)
	mux.HandleFunc("/translate/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck)
	mux.HandleFunc("/info", handlers.Info)

	log.Printf("listening on %s\n", addr) // <4>

	log.Fatal(http.ListenAndServe(addr, mux)) // <5>
}
