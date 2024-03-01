package main

type Range uint16

const (
	RangeMask = 0xff00

	RangeA uint16 = 0x1
	RangeB uint16 = 0x2
	RangeC uint16 = 0x4
	RangeD uint16 = 0x8
	RangeE uint16 = 0x10
	RangeF uint16 = 0x20
	RangeG uint16 = 0x40

	RangeVoltageDc  Range = 0x16 << 8
	RangeVoltageDcB       = RangeVoltageDc | Range(RangeB) // mV DC [000.0] E-1
	RangeVoltageDcC       = RangeVoltageDc | Range(RangeC) //  V DC [0.000] E-3
	RangeVoltageDcD       = RangeVoltageDc | Range(RangeD) //  V DC [00.00] E-2
	RangeVoltageDcE       = RangeVoltageDc | Range(RangeE) //  V DC [000.0] E-1
	RangeVoltageDcF       = RangeVoltageDc | Range(RangeF) //  V DC [0000.] E-0

	RangeVoltageAc  Range = 0x15 << 8
	RangeVoltageAcC       = RangeVoltageAc | Range(RangeC) // V AC [0.000] E-3
	RangeVoltageAcD       = RangeVoltageAc | Range(RangeD) // V AC [00.00] E-2
	RangeVoltageAcE       = RangeVoltageAc | Range(RangeE) // V AC [000.0] E-1
	RangeVoltageAcF       = RangeVoltageAc | Range(RangeF) // V AC [0000.] E-0

	RangeCurrent  Range = 0x18 << 8
	RangeCurrentF       = RangeCurrent | Range(RangeF) // A [0.000] E-3
	RangeCurrentG       = RangeCurrent | Range(RangeG) // A [00.00] E-2

	RangeCurrentMilli  Range = 0x19 << 8
	RangeCurrentMilliD       = RangeCurrentMilli | Range(RangeD) // mA [00.00] E-2
	RangeCurrentMilliE       = RangeCurrentMilli | Range(RangeE) // mA [000.0] E-1
	RangeCurrentMilliF       = RangeCurrentMilli | Range(RangeF) // mA [0.000] E-3

	RangeCurrentMicro  Range = 0x1a << 8
	RangeCurrentMicroB       = RangeCurrentMicro | Range(RangeB) // uA [000.0] E-1
	RangeCurrentMicroC       = RangeCurrentMicro | Range(RangeC) // uA [0000.] E-0

	RangeResistance  Range = 0x1d << 8
	RangeResistanceA       = RangeResistance | Range(RangeA) //  Ohm [000.0] E-1
	RangeResistanceB       = RangeResistance | Range(RangeB) // kOhm [0.000] E-3
	RangeResistanceC       = RangeResistance | Range(RangeC) // kOhm [00.00] E-2
	RangeResistanceD       = RangeResistance | Range(RangeD) // kOhm [000.0] E-1
	RangeResistanceE       = RangeResistance | Range(RangeE) // MOhm [0.000] E-3
	RangeResistanceF       = RangeResistance | Range(RangeF) // MOhm [00.00] E-2

	RangeContinuity  Range = 0x1b << 8
	RangeContinuityA       = RangeContinuity | Range(RangeA) // Ohm [000.0] E-1
	RangeContinuityC       = RangeContinuity | Range(RangeC) // Only for OL

	RangeDiode  Range = 0x1c << 8
	RangeDiodeC       = RangeDiode | Range(RangeC) // V [0.000] E-3

	RangeSquareWave Range = 0x3 << 8
)

func (r Range) Attributes() (a Attributes) {
	a = Attributes{
		Mode:      ModeUnknown,
		Polarity:  PolarityNone,
		Unit:      UnitNone,
		Recorded:  true,
		Precision: 0,
	}

	switch r & RangeMask {
	case RangeVoltageDc:
		a.Mode = ModeVoltage
		a.Polarity = PolarityDC
		a.Unit = UnitVolt
	case RangeVoltageAc:
		a.Mode = ModeVoltage
		a.Polarity = PolarityAC
		a.Unit = UnitVolt
	case RangeCurrent:
		a.Mode = ModeCurrent
		a.Unit = UnitAmpere
	case RangeCurrentMilli:
		a.Mode = ModeCurrent
		a.Unit = UnitAmpereMilli
	case RangeCurrentMicro:
		a.Mode = ModeCurrent
		a.Unit = UnitAmpereMicro
	case RangeResistance:
		a.Mode = ModeResistance
		a.Unit = UnitOhm
	case RangeContinuity:
		a.Mode = ModeContinuity
		a.Unit = UnitOhm
	case RangeDiode:
		a.Mode = ModeDiode
		a.Unit = UnitVolt
	case RangeSquareWave:
		a.Mode = ModeSquareWave
	}

	switch r {
	case RangeVoltageDcB:
		a.Unit = UnitVoltMilli
		a.Precision = 1
	case RangeVoltageDcC:
		a.Precision = 3
	case RangeVoltageDcD:
		a.Precision = 2
	case RangeVoltageDcE:
		a.Precision = 1
	case RangeVoltageDcF:
		a.Precision = 0
	case RangeVoltageAcC:
		a.Precision = 3
	case RangeVoltageAcD:
		a.Precision = 2
	case RangeVoltageAcE:
		a.Precision = 1
	case RangeVoltageAcF:
		a.Precision = 0
	case RangeCurrentF:
		a.Precision = 3
	case RangeCurrentG:
		a.Precision = 2
	case RangeCurrentMilliD:
		a.Precision = 2
	case RangeCurrentMilliE:
		a.Precision = 1
	case RangeCurrentMilliF:
		a.Precision = 3
	case RangeCurrentMicroB:
		a.Precision = 1
	case RangeCurrentMicroC:
		a.Precision = 0
	case RangeResistanceA:
		a.Precision = 1
	case RangeResistanceB:
		a.Unit = UnitOhmKilo
		a.Precision = 3
	case RangeResistanceC:
		a.Unit = UnitOhmKilo
		a.Precision = 2
	case RangeResistanceD:
		a.Unit = UnitOhmKilo
		a.Precision = 1
	case RangeResistanceE:
		a.Unit = UnitOhmMega
		a.Precision = 3
	case RangeResistanceF:
		a.Unit = UnitOhmMega
		a.Precision = 2
	case RangeContinuityA, RangeContinuityC:
		a.Precision = 1
	case RangeDiodeC:
		a.Precision = 3
	default:
		a.Recorded = false
	}
	return
}
