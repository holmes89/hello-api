// Package faas is used to house functions as a service methods.
package faas

import (
	"net/http"

	"github.com/holmes89/hello-api/handlers/rest"
)

// Translate is a handler for Google Cloud functions to use our built
// in translate handler.
func Translate(w http.ResponseWriter, r *http.Request) {
	rest.TranslateHandler(w, r)
}
