package main

import (
	"time"
)
import "go.bug.st/serial"

type Multimeter struct {
	port    string
	timeout time.Duration
	mode    serial.Mode
	serial  serial.Port
}

func NewMultimeter(port string, bitrate int, timeout time.Duration) (m *Multimeter) {
	m = &Multimeter{
		port:    port,
		timeout: timeout,
		mode: serial.Mode{
			BaudRate: bitrate,
			DataBits: 8,
			Parity:   serial.NoParity,
			StopBits: serial.OneStopBit,
			InitialStatusBits: &serial.ModemOutputBits{
				RTS: false,
				DTR: false,
			},
		},
	}
	return
}

func (m *Multimeter) Connect() (err error) {
	if m.serial != nil {
		return
	}

	if m.serial, err = serial.Open(m.port, &m.mode); err == nil {
		err = m.serial.SetReadTimeout(m.timeout)
	}
	return
}

func (m *Multimeter) Disconnect() (err error) {
	if m.serial == nil {
		return
	}

	err = m.serial.Close()
	m.serial = nil
	return
}
