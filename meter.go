package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"
	"sync"
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

func (m *Multimeter) Listen(output *csv.Writer) (read func() Reading, stop func(), err error) {
	if err = m.Connect(); err != nil {
		return
	}

	reading := &Reading{}
	readMutex := &sync.RWMutex{}
	read = func() Reading {
		readMutex.RLock()
		defer readMutex.RUnlock()
		return *reading
	}

	stopReq := make(chan struct{})
	stopOk := make(chan struct{})
	stop = func() {
		stopReq <- struct{}{}
		<-stopOk
	}

	go func() {
		defer func() {
			if err := m.Disconnect(); err != nil {
				log.Println(err)
			}

			log.Println("Stopped multimeter")
			stopOk <- struct{}{}
			close(stopOk)
		}()
		log.Printf("Start reading multimeter on %s at %d bps", m.port, m.mode.BaudRate)

		for {
			select {
			case <-stopReq:
				log.Println("Stopping multimeter")
				return
			default:
			}

			var ok bool
			if ok, err = m.Synchronize(); err != nil {
				panic(err)
			} else if !ok {
				log.Println("Failed to synchronize multimeter")
				continue
			}

			var r Reading
			if r, err = m.Receive(); err != nil {
				panic(err)
			}

			switch {
			case !r.Valid:
				log.Printf("Invalid multimeter packet: %+v", r)
				continue
			case r.Recorded:
				log.Printf("Multimeter packet (%s): %.3f%s%s %.1f%%", r.Mode.Translation(), r.Absolute, r.Unit, r.Polarity, r.Relative*100)
			default:
				log.Printf("Multimeter packet (%s): %s%s", r.Mode.Translation(), r.Unit, r.Polarity)
			}

			func() {
				readMutex.Lock()
				defer readMutex.Unlock()
				*reading = r
			}()

			// date, mode, rel, abs, unit, polarity
			if output != nil && r.Valid && r.Recorded && !r.Overload {
				if err = output.Write([]string{
					r.Received.String(),
					Translations[language][r.Mode.Translation()],
					strconv.FormatFloat(r.Relative, 'g', 3, 64),
					strconv.FormatFloat(r.Absolute, 'f', r.Precision, 64),
					string(r.Unit),
					string(r.Polarity),
				}); err != nil {
					log.Printf("Could not write CSV: %v", err)
					continue
				}

				output.Flush()
				if err = output.Error(); err != nil {
					log.Printf("Could not flush CSV: %v", err)
					continue
				}
			}
		}
	}()
	return
}

func (m *Multimeter) Synchronize() (ok bool, err error) {
	if m.serial == nil {
		err = fmt.Errorf("the port is not open")
	}

	if err = m.serial.ResetInputBuffer(); err != nil {
		return
	}

	var buf = make([]byte, 1)
	var tail bool
	var retry int
	m.sync = false
	for retry = 0; retry < 11; retry++ {
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
			m.sync = true
			return
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
	}

	var buf = make([]byte, 8)
	if _, err = io.ReadFull(m.serial, buf); err != nil {
		return
	}
	r.Received = time.Now()

	var actualChecksum uint16
	var expectedChecksum = (uint16(buf[6]) << 8) | uint16(buf[7])
	for i := 0; i <= 5; i++ {
		actualChecksum += uint16(buf[i])
	}
	if r.Valid = actualChecksum == expectedChecksum; !r.Valid {
		return
	}

	r.Attributes = Range(uint16(buf[1])<<8 | uint16(buf[2])).Attributes()
	if r.Recorded && r.Maximum > r.Minimum {
		raw := int16(uint16(buf[4])<<8 | uint16(buf[5]))
		r.Relative = math.Max(math.Min(float64(raw-r.Minimum)/float64(r.Maximum-r.Minimum), 1), 0)
		if raw >= r.Minimum && raw <= r.Maximum {
			r.Absolute = float64(raw) * math.Pow10(-r.Precision)
		} else {
			r.Overload = true
		}
	}
	return
}
