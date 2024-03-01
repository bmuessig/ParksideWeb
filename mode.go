package main

// Mode that the reading was taken in.
type Mode uint

const (
	ModeUnknown Mode = iota
	ModeVoltage
	ModeCurrent
	ModeResistance
	ModeContinuity
	ModeDiode
	ModeSquareWave
)

func (m Mode) String(l Language) (s string, ok bool) {
	var t Translation
	switch m {
	case ModeVoltage:
		t = TranslationVoltage
	case ModeCurrent:
		t = TranslationCurrent
	case ModeResistance:
		t = TranslationResistance
	case ModeContinuity:
		t = TranslationContinuity
	case ModeDiode:
		t = TranslationDiode
	case ModeSquareWave:
		t = TranslationSquareWave
	default:
		t = TranslationUnknown
	}

	s, ok = Translations[l][t]
	return
}
