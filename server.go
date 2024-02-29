package main

import (
	"embed"
	"encoding/json"
	"math/rand"
	"net/http"
)

//go:embed main.html style.css casper.woff favicon.ico
var appFs embed.FS

func Serve(writer http.ResponseWriter, request *http.Request) {
	// Allow access from anywhere
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	writer.Header().Set("Access-Control-Request-Method", "*")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept")
	if request.Method == http.MethodOptions {
		writer.Header().Set("Allow", "OPTIONS, GET")
		writer.WriteHeader(http.StatusNoContent)
		return
	} else if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	switch request.URL.Path {
	case "", "/":
		request.URL.Path = "main.html"

	case "config.json", "/config.json":
		writer.Header().Set("Content-Type", "application/json")

		b, _ := json.Marshal(struct {
			Version    string   `json:"version"`
			Language   Language `json:"language"`
			Whitelabel bool     `json:"whitelabel"`
			Port       string   `json:"port"`
		}{
			Version:    "1.0",
			Language:   LanguageEnglish,
			Whitelabel: true,
			Port:       "/dev/ttyUSB0",
		})

		writer.WriteHeader(http.StatusOK)
		writer.Write(b)

		// todo
		return

	case "data.json", "/data.json":
		writer.Header().Set("Content-Type", "application/json")

		val := (float64)(rand.Int()%10000) / 10
		if (rand.Int() % 2) == 0 {
			val = -val
		}

		b, _ := json.Marshal(struct {
			Valid bool    `json:"valid"`
			Time  int64   `json:"time"`
			Value float64 `json:"value"`
			Unit  string  `json:"unit"`
			Mode  string  `json:"mode"`
		}{
			Valid: true,
			Time:  0,
			Value: val,
			Unit:  "mV\nDC",
			Mode:  Translations[LanguageGerman][TranslationVoltage],
		})

		writer.WriteHeader(http.StatusOK)
		writer.Write(b)

		// todo
		return
	}

	http.FileServer(http.FS(appFs)).ServeHTTP(writer, request)
}
