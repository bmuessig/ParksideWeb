package main

// Range that the value was taken in.
type Range byte

const (
	// 0000 0000 (observed during startup)
	RangeNone Range = 0x0

	// 0000 0001
	RangeA Range = 0x1

	// 0000 0010
	RangeB = 0x2

	// 0000 0100
	RangeC = 0x4

	// 0000 1000
	RangeD = 0x8

	// 0001 0000
	RangeE = 0x10

	// 0010 0000
	RangeF = 0x20

	// 0100 0000
	RangeG = 0x40

	// 1000 0000
	RangeH = 0x80
)
