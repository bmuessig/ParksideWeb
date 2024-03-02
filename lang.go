package main

type Language string

const (
	LanguageEnglish    Language = "en"
	LanguageGerman     Language = "de"
	LanguagePortuguese Language = "pt"
	LanguageFrench     Language = "fr"
)

type Translation string

const (
	TranslationUnknown    Translation = "unknown"
	TranslationSquareWave Translation = "squareWave"
	TranslationVoltage    Translation = "voltage"
	TranslationResistance Translation = "resistance"
	TranslationDiode      Translation = "diode"
	TranslationContinuity Translation = "continuity"
	TranslationCurrent    Translation = "current"
	TranslationDate       Translation = "date"
	TranslationMode       Translation = "mode"
	TranslationRelative   Translation = "relative"
	TranslationAbsolute   Translation = "absolute"
	TranslationUnit       Translation = "unit"
	TranslationPolarity   Translation = "polarity"
)

var Translations = map[Language]map[Translation]string{
	LanguageEnglish: {
		TranslationUnknown:    "Unknown",
		TranslationSquareWave: "Square wave",
		TranslationVoltage:    "Voltage",
		TranslationResistance: "Resistance",
		TranslationDiode:      "Diode",
		TranslationContinuity: "Continuity",
		TranslationCurrent:    "Current",
		TranslationDate:       "Date",
		TranslationMode:       "Mode",
		TranslationRelative:   "Relative",
		TranslationAbsolute:   "Absolute",
		TranslationUnit:       "Unit",
		TranslationPolarity:   "Polarity",
	},
	LanguageGerman: {
		TranslationUnknown:    "Unbekannt",
		TranslationSquareWave: "Rechtecksignal",
		TranslationVoltage:    "Spannung",
		TranslationResistance: "Widerstand",
		TranslationDiode:      "Diode",
		TranslationContinuity: "Kontinuität",
		TranslationCurrent:    "Strom",
		TranslationDate:       "Datum",
		TranslationMode:       "Modus",
		TranslationRelative:   "Relativ",
		TranslationAbsolute:   "Absolut",
		TranslationUnit:       "Einheit",
		TranslationPolarity:   "Polarität",
	},
	LanguagePortuguese: {
		TranslationUnknown:    "Desconhecido",
		TranslationSquareWave: "Onda quadrada",
		TranslationVoltage:    "Tensão",
		TranslationResistance: "Resistência",
		TranslationDiode:      "Díodo",
		TranslationContinuity: "Continuidade",
		TranslationCurrent:    "Corrente",
		TranslationDate:       "Data",
		TranslationMode:       "Modo",
		TranslationRelative:   "Relativo",
		TranslationAbsolute:   "Absoluto",
		TranslationUnit:       "Unidade",
		TranslationPolarity:   "Polaridade",
	},
	LanguageFrench: {
		TranslationUnknown:    "Inconnu",
		TranslationSquareWave: "Onde carrée",
		TranslationVoltage:    "Tension",
		TranslationResistance: "Résistance",
		TranslationDiode:      "Diode",
		TranslationContinuity: "Continuité",
		TranslationCurrent:    "Courant",
		TranslationDate:       "Date",
		TranslationMode:       "Mode",
		TranslationRelative:   "Relatif",
		TranslationAbsolute:   "Absolu",
		TranslationUnit:       "Unité",
		TranslationPolarity:   "Polarité",
	},
}
