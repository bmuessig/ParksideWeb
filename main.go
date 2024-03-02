package main

import (
	"context"
	"encoding/csv"
	"errors"
	"github.com/pkg/browser"
	"go.bug.st/serial"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)
import (
	"fmt"
	"net/http"
)

func main() {
	ParseFlags()

	if serialEnumerate {
		ports, err := serial.GetPortsList()
		if err != nil {
			log.Printf("Could not get ports list: %v", err)
			return
		}

		for _, port := range ports {
			fmt.Println(port)
		}
		return
	}

	log.Printf("ParksideWeb v%s", version)
	log.Println("(c) Benedikt Muessig, 2024")
	log.Println("https://github.com/bmuessig/ParksideWeb")
	if serialPort == "" {
		log.Println("No serial port specified (e.g. -s COM1 or /dev/ttyUSB0)")
		return
	}

	var durationExpired <-chan time.Time
	if duration != 0 {
		durationExpired = time.NewTimer(duration).C
	}

	var csvWriter *csv.Writer
	if csvOutput {
		csvWriter = csv.NewWriter(os.Stdout)
		if err := csvWriter.Write([]string{
			Translations[language][TranslationDate],
			Translations[language][TranslationMode],
			Translations[language][TranslationRelative],
			Translations[language][TranslationAbsolute],
			Translations[language][TranslationUnit],
			Translations[language][TranslationPolarity],
		}); err != nil {
			log.Printf("Could not write CSV: %v", err)
			return
		}
	}

	multimeter := NewMultimeter(serialPort, serialBitrate, timeout)
	var err error
	var getReading func() Reading
	var stopReading func()
	if getReading, stopReading, err = multimeter.Listen(csvWriter); err != nil {
		log.Printf("Could not start multimeter: %v", err)
		return
	}

	if !serverDisable {
		server := &http.Server{
			Handler: NewHandler(getReading),
		}
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
		if err != nil {
			panic(err)
		}

		port := listener.Addr().(*net.TCPAddr).Port
		url := fmt.Sprintf("http://localhost:%d", port)
		log.Printf("Starting server on %s", url)
		go func() {
			if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
				log.Printf("Could start HTTP server: %v", err)
			}
			log.Println("Stopped serving new HTTP requests")
		}()

		if openBrowser {
			browser.Stdout = log.Writer()
			browser.Stderr = browser.Stdout
			if err = browser.OpenURL(url); err != nil {
				log.Printf("Could not open URL in browser: %v", err)
			}
		}

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sigChan:
		case <-durationExpired:
		}

		shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownRelease()

		if err = server.Shutdown(shutdownCtx); err != nil {
			log.Printf("Could not stop HTTP server: %v", err)
			if err = server.Close(); err != nil {
				log.Printf("Could not force-stop HTTP server: %v", err)
			}
		}
	} else {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sigChan:
		case <-durationExpired:
		}
	}

	stopReading()
	log.Println("Shutdown complete")
}
