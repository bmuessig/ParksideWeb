package main

import "time"

const (
	version = "1.0"
	timeout = time.Second
)

var (
	openBrowser   bool
	language      Language
	serialPort    string
	serialBitrate int
	serverPort    int
	csvOutput     bool
	duration      time.Duration
)
