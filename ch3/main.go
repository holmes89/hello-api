package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	cfg, err := LoadConfiguration()
	if err != nil {
		log.Fatal(err)
	}

	addr := cfg.Port

	log.Printf("using default language: %s", cfg.DefaultLanguage)
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		resp := Resp{
			Language:    cfg.DefaultLanguage,
			Translation: langs[cfg.DefaultLanguage],
		}
		if err := enc.Encode(resp); err != nil {
			panic("unable to encode response")
		}
	})

	fmt.Printf("listening on %s\n", addr)

	log.Fatal(http.ListenAndServe(addr, mux))
}

var langs = map[string]string{
	"English": "Hello",
	"Finnish": "Hei",
}

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

type configuration struct {
	DefaultLanguage string `json:"default_language"`
	Port            string `json:"port"`
}

func LoadConfiguration() (configuration, error) {
	cfgfilePtr := flag.String("config_file", "", "load configurations from a file")
	portPtr := flag.String("port", "", "set port")

	flag.Parse()

	cfg := defaultConfig
	cfg.LoadFromEnv()

	if cfgfilePtr != nil && *cfgfilePtr != "" {
		if err := cfg.LoadFromJSON(*cfgfilePtr); err != nil {
			return cfg, fmt.Errorf("unable to load configuration from json: %s\n", *cfgfilePtr)
		}
	}

	if portPtr != nil && *portPtr != "" {
		cfg.Port = *portPtr
	}
	cfg.ParsePort()
	return cfg, nil
}

func (c *configuration) ParsePort() {
	if c.Port[0] != ':' {
		c.Port = ":" + c.Port
	}
}

var defaultConfig = configuration{
	DefaultLanguage: "English",
	Port:            ":8080",
}

func (c *configuration) LoadFromEnv() {
	if lang := os.Getenv("DEFAULT_LANGUAGE"); lang != "" {
		c.DefaultLanguage = lang
	}
	if port := os.Getenv("PORT"); port != "" {
		c.Port = port
	}
}

func (c *configuration) LoadFromJSON(path string) error {
	log.Printf("loading configuration from file: %s\n", path)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("unable to load file: %s\n", err.Error())
		return errors.New("unable to load configuration")
	}
	if err := json.Unmarshal(b, c); err != nil {
		log.Printf("unable to parse file: %s\n", err.Error())
		return errors.New("unable to load configuration")
	}
	return nil
}
