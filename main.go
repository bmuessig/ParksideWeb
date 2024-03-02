package main

import (
	"context"
	"errors"
	"github.com/pkg/browser"
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

/*
	-o open browser (default false)
	-l language (de, en, pt, fr?)
	-p serial port
	-b serial bitrate (default 2400)
	-s override HTTP port (default 0 auto)
	-w white label mode (default false)
	-c dump csv to stdout and don't run server (default false)
	-d minimum duration to run the acquisition for (default 0 unlimited)
	-h help
*/

func main() {
	serialPort = "/dev/cu.usbmodem143301"
	serialBitrate = 2400
	language = LanguageEnglish
	openBrowser = true

	server := &http.Server{
		Addr: ":8080",
	}

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		panic(err)
	}

	port := listener.Addr().(*net.TCPAddr).Port
	log.Println("Using port:", port)

	multimeter := NewMultimeter(serialPort, serialBitrate, timeout)
	var getReading func() Reading
	var stopReading func()
	if getReading, stopReading, err = multimeter.Listen(); err != nil {
		return
	}

	server.Handler = NewHandler(getReading)
	go func() {
		if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Could start HTTP server: %v", err)
		}
		log.Println("Stopped serving new HTTP requests")
	}()

	if openBrowser {
		browser.Stdout = log.Writer()
		browser.Stderr = browser.Stdout
		browser.OpenURL(fmt.Sprintf("http://localhost:%d", port))
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("Could not stop HTTP server: %v", err)
		if err := server.Close(); err != nil {
			log.Printf("Could not force-stop HTTP server: %v", err)
		}
	}

	stopReading()
	log.Println("Shutdown complete")
}
