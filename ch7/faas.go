package faas

import (
	"net/http"

	"github.com/holmes89/hello-api/handlers/rest"

	"github.com/holmes89/hello-api/translation"
)

func Translate(w http.ResponseWriter, r *http.Request) {
	translationService := translation.NewStaticService()
	translateHandler := rest.NewTranslateHandler(translationService)

	translateHandler.TranslateHandler(w, r)
}
