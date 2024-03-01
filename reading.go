package main

import "time"

type Polarity string

const (
	PolarityNone Polarity = ""
	PolarityDC   Polarity = "DC"
	PolarityAC   Polarity = "AC"
)

type Unit string

const (
	UnitNone        Unit = ""
	UnitAmpere      Unit = "A"
	UnitAmpereMilli Unit = "mA"
	UnitAmpereMicro Unit = "µA"
	UnitVolt        Unit = "V"
	UnitVoltMilli   Unit = "mV"
	UnitOhm         Unit = "Ω"
	UnitOhmKilo     Unit = "kΩ"
	UnitOhmMega     Unit = "MΩ"
)

type Attributes struct {
	Mode      Mode
	Polarity  Polarity
	Unit      Unit
	Recorded  bool
	Precision int
}

type Reading struct {
	Received time.Time
	Valid    bool
	Attributes
	Value float64
}
