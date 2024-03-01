package main

import (
	"fmt"
	"io"
	"math"
	"time"
)
import "go.bug.st/serial"

type Multimeter struct {
	port    string
	timeout time.Duration
	mode    serial.Mode
	serial  serial.Port
	sync    bool
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

func (m *Multimeter) Synchronize() (ok bool, err error) {
	if m.serial == nil {
		err = fmt.Errorf("the port is not open")
	}

	var buf = make([]byte, 1)
	var tail bool
	var retry int
	for retry = 10; retry > 0; retry-- {
		n := 0
		n, err = m.serial.Read(buf)
		switch {
		case n != 1:
			continue
		case err != nil:
			return
		case buf[0] == 0xdc:
			tail = true
		case buf[0] == 0xba:
			if !tail {
				continue
			}
			ok = true
			break
		default:
			tail = false
		}
	}
	return
}

func (m *Multimeter) Receive() (r Reading, err error) {
	switch {
	case m.serial == nil:
		err = fmt.Errorf("the port is not open")
		return
	case !m.sync:
		err = fmt.Errorf("the connection is not synchronized")
		return
	default:
		m.sync = false
		r.Received = time.Now()
	}

	var buf = make([]byte, 8)
	if _, err = io.ReadFull(m.serial, buf); err != nil {
		return
	}

	var actualChecksum uint16
	var expectedChecksum = (uint16(buf[6]) << 8) | uint16(buf[7])
	for i := 0; i <= 5; i++ {
		actualChecksum += uint16(buf[i])
	}
	if actualChecksum != expectedChecksum {
		return
	}

	r.Attributes = Range(uint16(buf[1])<<8 | uint16(buf[2])).Attributes()
	if r.Recorded {
		r.Value = float64(int16(uint16(buf[4])<<8|uint16(buf[5]))) * math.Pow10(-r.Precision)
	}
	return
}
