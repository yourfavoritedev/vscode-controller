package themes

// AllThemes is a map of all default themes
var AllThemes = map[string]string{}

func init() {
	AllThemes = map[string]string{
		"Default Dark+":       "Default Dark+",
		"Abyss":               "Abyss",
		"Quiet Light":         "Quiet Light",
		"Default Light+":      "Default Light+",
		"Monokai":             "Monokai",
		"Red":                 "Red",
		"Tomorrow Night Blue": "Tomorrow Night Blue",
	}
}
