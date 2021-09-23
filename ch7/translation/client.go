package translation

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

var _ HelloClient = &APIClient{}

type APIClient struct {
	endpoint string
}

func NewHelloClient(endpoint string) *APIClient {
	return &APIClient{
		endpoint: endpoint,
	}
}

func (c *APIClient) Translate(word, language string) (string, error) {
	req := map[string]interface{}{
		"word":     word,
		"language": language,
	}
	b, err := json.Marshal(req)
	if err != nil {
		return "", errors.New("unable to encode msg")
	}

	resp, err := http.Post(c.endpoint, "application/json", bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		return "", errors.New("call to api failed")
	}

	if resp.StatusCode == http.StatusNotFound {
		return "", nil
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return "", errors.New("error in api")
	}

	b, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return "", errors.New("unable to decode message")
	}

	return m["translation"].(string), nil
}
