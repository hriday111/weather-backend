package util

var translations = map[string]map[string]string{
	"en": {
		"with_precipitation":    "with precipitation",
		"without_precipitation": "without precipitation",
		"monday":                "Monday",
		"tuesday":               "Tuesday",
		"wednesday":             "Wednesday",
		"thursday":              "Thursday",
		"friday":                "Friday",
		"saturday":              "Saturday",
		"sunday":                "Sunday",
	},
	"pl": {
		"with_precipitation":    "z opadami",
		"without_precipitation": "bez opadów",
		"monday":                "Poniedziałek",
		"tuesday":               "Wtorek",
		"wednesday":             "Środa",
		"thursday":              "Czwartek",
		"friday":                "Piątek",
		"saturday":              "Sobota",
		"sunday":                "Niedziela",
	},
}

func Translate(key, lang string) string {
	if langMap, ok := translations[lang]; ok {
		if val, ok := langMap[key]; ok {
			return val
		}
	}
	// fallback to English
	return translations["en"][key]
}
