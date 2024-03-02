package main

import (
	"flag"
	"time"
)

const (
	version = "1.0"
	timeout = time.Second
)

var (
	openBrowser   bool
	serverPort    int
	serverDisable bool

	serialPort      string
	serialBitrate   int
	serialEnumerate bool

	language  Language
	csvOutput bool
	duration  time.Duration
)

func ParseFlags() {
	flag.BoolVar(&openBrowser, "o", false, "open the live display in a browser")
	flag.IntVar(&serverPort, "s", 0, "choose a static port number over a free port")
	flag.BoolVar(&serverDisable, "n", false, "disable the HTTP server")

	flag.StringVar(&serialPort, "p", "", "name or path of the serial port")
	flag.IntVar(&serialBitrate, "b", 2400, "override the default serial port bitrate")
	flag.BoolVar(&serialEnumerate, "e", false, "enumerate all serial ports then exit")

	l := flag.String("l", "en", "language for CSV and display (en,de,pt,fr)")
	flag.BoolVar(&csvOutput, "c", false, "enable CSV output on standard output")
	flag.DurationVar(&duration, "d", 0, "choose a acquisition duration over unlimited")
	flag.Parse()

	language = LanguageEnglish
	if _, ok := Translations[Language(*l)]; ok {
		language = Language(*l)
	}
}
