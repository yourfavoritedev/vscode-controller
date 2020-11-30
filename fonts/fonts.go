package fonts

// AllFonts is a map of all default fonts
var AllFonts = map[string]string{}

func init() {
	AllFonts = map[string]string{
		"Consolas":    "Consolas",
		"Calibri":     "Calibri",
		"Arial":       "Arial",
		"Bahnschrift": "Bahnschrift",
		"Cambria":     "Cambria",
		"Candara":     "Candara",
		"Corbel":      "Corbel",
		"Impact":      "Impact",
	}
}
