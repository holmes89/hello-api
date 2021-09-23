package translation_test

import (
	"testing"

	"github.com/holmes89/hello-api/translation"
)

func TestTranslate(t *testing.T) {
	// Arrange
	tt := []struct { // <1>
		Word        string
		Language    string
		Translation string
	}{
		{ //<2>
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{ // <1>
			Word:        "bye",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
		{ // <1>
			Word:        "bye",
			Language:    "german",
			Translation: "",
		},
		{
			Word:        "hello",
			Language:    "German", // <1>
			Translation: "hallo",
		},
		{
			Word:        "Hello", // <2>
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello ", // <3>
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "french",
			Translation: "bonjour",
		},
	}
	underTest := translation.NewStaticService()
	for _, test := range tt { // <3>
		// Act
		res := underTest.Translate(test.Word, test.Language) // <4>

		// Assert
		if res != test.Translation { // <5>
			t.Errorf(
				`expected "%s" to be "%s" from "%s" but received "%s"`,
				test.Word, test.Language, test.Translation, res)
		}
	}
}
