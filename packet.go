package main

import "time"

// Packet of data received from the multimeter.
type Packet struct {
	// Mode that the packet was taken in.
	Mode Mode

	// Range that the value was taken in.
	Range Range

	// Value of the packet.
	Value int16

	// Valid indicates, whether the packet checksum of packets 2-7 checked out.
	Valid bool

	// Unknown data (2 byte) in each packet that has not yet been decoded.
	Unknown []byte

	// Received is the time at which the packet was received.
	Received time.Time
}

type Polarity string

const (
	PolarityNone Polarity = ""
	PolarityDC   Polarity = "DC"
	PolarityAC   Polarity = "AC"
)

type Unit string

const (
	UnitAmpere      string = "A"
	UnitAmpereMilli string = "mA"
	UnitAmpereMicro string = "µA"
	UnitVolt               = "V"
	UnitVoltMilli          = "mV"
)

type Reading struct {
	Received  time.Time
	Valid     bool
	Mode      Mode // TODO
	Unit      Unit
	Polarity  Polarity
	Value     float64
	Precision uint
}
