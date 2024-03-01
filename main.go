package main

import (
	"github.com/pkg/browser"
	"log"
	"net"
)
import (
	"fmt"
	"net/http"
	"sync"
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

var reading = &Reading{}
var mutex = &sync.RWMutex{}

func main() {
	serialPort = "/dev/cu.usbmodem143301"
	serialBitrate = 2400
	language = LanguageEnglish
	openBrowser = true

	go func() {
		m := NewMultimeter(serialPort, serialBitrate, timeout)
		err := m.Connect()
		if err != nil {
			panic(err)
		}
		defer m.Disconnect()

		for {
			var ok bool
			if ok, err = m.Synchronize(); err != nil {
				panic(err)
			} else if !ok {
				fmt.Println("Sync failed")
				continue
			}

			var r Reading
			if r, err = m.Receive(); err != nil {
				panic(err)
			}

			switch {
			case !r.Valid:
				fmt.Printf("%v: Invalid packet\n", r.Received)
				continue
			case r.Recorded:
				mode, _ := r.Mode.String(LanguageEnglish)
				fmt.Printf("%v (%s): %f%s%s %f%%\n", r.Received, mode, r.Absolute, r.Unit, r.Polarity, r.Relative*100)
			default:
				mode, _ := r.Mode.String(LanguageEnglish)
				fmt.Printf("%v (%s): %s%s\n", r.Received, mode, r.Unit, r.Polarity)
			}

			func() {
				mutex.Lock()
				defer mutex.Unlock()

				*reading = r
			}()
		}
	}()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", serverPort))
	if err != nil {
		panic(err)
	}

	port := listener.Addr().(*net.TCPAddr).Port
	fmt.Println("Using port:", port)

	if openBrowser {
		browser.Stdout = log.Writer()
		browser.Stderr = browser.Stdout
		browser.OpenURL(fmt.Sprintf("http://localhost:%d", port))
	}

	http.HandleFunc("/", Serve)
	http.Serve(listener, nil)
}
