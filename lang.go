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
	},
	LanguageGerman: {
		TranslationUnknown:    "Unbekannt",
		TranslationSquareWave: "Rechtecksignal",
		TranslationVoltage:    "Spannung",
		TranslationResistance: "Widerstand",
		TranslationDiode:      "Diode",
		TranslationContinuity: "Kontinuität",
		TranslationCurrent:    "Strom",
	},
	LanguagePortuguese: {
		TranslationUnknown:    "Desconhecido",
		TranslationSquareWave: "Onda quadrada",
		TranslationVoltage:    "Tensão",
		TranslationResistance: "Resistência",
		TranslationDiode:      "Díodo",
		TranslationContinuity: "Continuidade",
		TranslationCurrent:    "Corrente",
	},
	LanguageFrench: {
		TranslationUnknown:    "Inconnu",
		TranslationSquareWave: "Onde carrée",
		TranslationVoltage:    "Tension",
		TranslationResistance: "Résistance",
		TranslationDiode:      "Diode",
		TranslationContinuity: "Continuité",
		TranslationCurrent:    "Courant",
	},
}
