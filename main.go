package main

import (
	"log"
	"net/http"
	"os"

	handler "github.com/holmes89/hello-api/handlers"
	"github.com/holmes89/hello-api/handlers/rest"
)

func main() {
	addr := os.Getenv("PORT")
	if addr == "" {
		addr = ":8080" // <1>
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/translate/hello", rest.TranslateHandler)
	mux.HandleFunc("/health", handler.HealthCheck)

	log.Printf("listening on %s\n", addr) // <4>

	log.Fatal(http.ListenAndServe(addr, mux)) // <5>
}
