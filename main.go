package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
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

var (
	openBrowser   bool
	language      string
	serialPort    string
	serialBitrate int
	serverPort    int
	whiteLabel    bool
	csvOutput     bool
	duration      time.Duration
)

var reading = &Reading{}
var mutex = &sync.RWMutex{}

func main() {
	go func() {
		m := NewMultimeter("/dev/cu.usbmodem143301", 2400, 1*time.Second)
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
				fmt.Printf("%v (%s): %f%s%s\n", r.Received, mode, r.Value, r.Unit, r.Polarity)
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

	http.HandleFunc("/", Serve)
	http.ListenAndServe(":8080", nil)
}
