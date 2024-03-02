package main

import (
	"embed"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

//go:embed main.html style.css casper.woff favicon.ico
var appFs embed.FS

type Handler struct {
	read func() Reading
}

func NewHandler(read func() Reading) *Handler {
	return &Handler{read: read}
}

func (s *Handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

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
			Version:  version,
			Language: language,
			Port:     serialPort,
		})

		writer.WriteHeader(http.StatusOK)
		writer.Write(b)

		// todo
		return

	case "data.json", "/data.json":
		writer.Header().Set("Content-Type", "application/json")

		reading := s.read()
		mode := Translations[language][reading.Mode.Translation()]
		stale := time.Now().Sub(reading.Received).Seconds() >= 2
		b, _ := json.Marshal(struct {
			Recorded bool    `json:"recorded"`
			Overload bool    `json:"overload"`
			Time     int64   `json:"time"`
			Value    float64 `json:"value"`
			Needle   int     `json:"needle"`
			Digits   int     `json:"digits"`
			Unit     string  `json:"unit"`
			Polarity string  `json:"polarity"`
			Mode     string  `json:"mode"`
		}{
			Recorded: reading.Valid && reading.Recorded && !stale,
			Overload: reading.Overload,
			Time:     reading.Received.UnixMilli(),
			Value:    reading.Absolute,
			Needle:   int(reading.Relative * 100),
			Digits:   reading.Precision,
			Unit:     string(reading.Unit),
			Polarity: string(reading.Polarity),
			Mode:     mode,
		})

		writer.WriteHeader(http.StatusOK)
		writer.Write(b)

		// todo
		return
	}

	http.FileServer(http.FS(appFs)).ServeHTTP(writer, request)
}
