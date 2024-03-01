package main

import (
	"embed"
	"encoding/json"
	"net/http"
	"time"
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
			Version  string   `json:"version"`
			Language Language `json:"language"`
			Port     string   `json:"port"`
		}{
			Version:  "1.0",
			Language: LanguageEnglish,
			Port:     "/dev/ttyUSB0",
		})

		writer.WriteHeader(http.StatusOK)
		writer.Write(b)

		// todo
		return

	case "data.json", "/data.json":
		writer.Header().Set("Content-Type", "application/json")

		b := func() []byte {
			mutex.RLock()
			defer mutex.RUnlock()

			mode, _ := reading.Mode.String(LanguageGerman)
			c, _ := json.Marshal(struct {
				Valid    bool    `json:"valid"`
				Time     int64   `json:"time"`
				Value    float64 `json:"value"`
				Needle   int     `json:"needle"`
				Digits   int     `json:"digits"`
				Unit     string  `json:"unit"`
				Polarity string  `json:"polarity"`
				Mode     string  `json:"mode"`
			}{
				Valid:    reading.Valid && time.Now().Sub(reading.Received).Seconds() < 2,
				Time:     reading.Received.UnixMilli(),
				Value:    reading.Absolute,
				Needle:   int(reading.Relative * 100),
				Digits:   reading.Precision,
				Unit:     string(reading.Unit),
				Polarity: string(reading.Polarity),
				Mode:     mode,
			})
			return c
		}()

		writer.WriteHeader(http.StatusOK)
		writer.Write(b)

		// todo
		return
	}

	http.FileServer(http.FS(appFs)).ServeHTTP(writer, request)
}
