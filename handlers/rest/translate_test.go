package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/holmes89/hello-api/handlers/rest"
)

type stubbedService struct{}

func (s *stubbedService) Translate(word string, language string) string {
	if word == "foo" {
		return "bar"
	}
	return ""
}

func TestTranslateAPI(t *testing.T) {

	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/translate/foo",
			StatusCode:          200,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "bar",
		},
		{
			Endpoint:            "/translate/foo?language=german",
			StatusCode:          200,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "bar",
		},
		{
			Endpoint:            "/translate/baz",
			StatusCode:          404,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}

	h := rest.NewTranslateHandler(&stubbedService{})
	handler := http.HandlerFunc(h.TranslateHandler)

	for _, test := range tt {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)

		handler.ServeHTTP(rr, req)

		if rr.Code != test.StatusCode {
			t.Errorf(`expected status %d but received %d`, test.StatusCode, rr.Code)
		}

		var resp rest.Resp
		_ = json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.Language != test.ExpectedLanguage {
			t.Errorf(`expected language "%s" but received "%s"`, test.ExpectedLanguage, resp.Language)
		}

		if resp.Translation != test.ExpectedTranslation {
			t.Errorf(`expected Translation "%s" but received "%s"`, test.ExpectedTranslation, resp.Translation)
		}
	}
}
