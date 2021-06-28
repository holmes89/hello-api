package main

import (
	"log"
	"net/http"

	"github.com/holmes89/hello-api/handlers/rest"
)

func main() {

	addr := ":8080" // <1>

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", rest.TranslateHandler)

	log.Printf("listening on %s\n", addr) // <4>

	log.Fatal(http.ListenAndServe(addr, mux)) // <5>
}
