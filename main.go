package main

import (
	"net/http"
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

func main() {
	http.HandleFunc("/", Serve)
	http.ListenAndServe(":8080", nil)
}
