package main

// Mode that the reading was taken in.
type Mode byte

const (
	// DC Volt
	VoltDC = 0x16

	// AC Volt
	VoltAC = 0x15

	// Microampere (AC or DC)
	AmpereMicro = 0x1a

	// Milliampere (AC or DC)
	AmpereMilli = 0x19

	// Ampere (AC or DC)
	Ampere = 0x18

	// Resistance in Ohms
	ResistanceOhm = 0x1d

	// Continuity in Ohms
	ContinuityOhm = 0x1b

	// Diode in Volts
	DiodeVolt = 0x1c

	// Squarewave mode (without readings)
	Squarewave = 0x3
)
