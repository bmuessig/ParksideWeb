package main

// Mode that the reading was taken in.
type Mode uint

const (
	ModeUnknown Mode = iota
	ModeBooting
	ModeVoltage
	ModeCurrent
	ModeResistance
	ModeContinuity
	ModeDiode
	ModeSquareWave
)
