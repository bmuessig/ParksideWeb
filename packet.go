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
