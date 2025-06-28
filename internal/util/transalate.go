package util

var translations = map[string]map[string]string{
	"en": {
		"with_precipitation":    "with precipitation",
		"without_precipitation": "without precipitation",
		"Monday":                "Monday",
		"Tuesday":               "Tuesday",
		"Wednesday":             "Wednesday",
		"Thursday":              "Thursday",
		"Friday":                "Friday",
		"Saturday":              "Saturday",
		"Sunday":                "Sunday",
	},
	"pl": {
		"with_precipitation":    "z opadami",
		"without_precipitation": "bez opadów",
		"Monday":                "Poniedziałek",
		"Tuesday":               "Wtorek",
		"Wednesday":             "Środa",
		"Thursday":              "Czwartek",
		"Friday":                "Piątek",
		"Saturday":              "Sobota",
		"Sunday":                "Niedziela",
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
