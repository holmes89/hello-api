package main

import (
	"log"
	"net/http"

	"github.com/holmes89/hello-api/config"
	"github.com/holmes89/hello-api/handlers"
	"github.com/holmes89/hello-api/handlers/rest"
	"github.com/holmes89/hello-api/translation"
)

func main() {

	cfg := config.LoadConfiguration()
	addr := cfg.Port

	mux := http.NewServeMux()

	var translationService rest.Translator
	translationService = translation.NewStaticService()
	if cfg.LegacyEndpoint != "" {
		log.Printf("creating external translation client: %s", cfg.LegacyEndpoint)
		client := translation.NewHelloClient(cfg.LegacyEndpoint)
		translationService = translation.NewRemoteService(client)
	}

	translateHandler := rest.NewTranslateHandler(translationService)
	mux.HandleFunc("/translate/hello", translateHandler.TranslateHandler)
	mux.HandleFunc("/health", handlers.HealthCheck) // <1>

	log.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}
