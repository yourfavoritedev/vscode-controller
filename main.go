package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	components "github.com/yourfavoritedev/vscode-controller/components"
)

func init() {
	// initializes ui
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
}

func main() {
	defer ui.Close()
	// define channel to share key-events between components
	eventChannel := make(chan string)

	themeShuffler, setThemeShufflerActive := components.ThemeShuffler(eventChannel)
	fontFamilyShuffler, setFontFamilyShufflerActive := components.FontShuffler(eventChannel)
	fontSizeSetter, setFontSetterActive := components.FontSizeSetter(eventChannel)

	// render the initial ui
	ui.Render(
		themeShuffler,
		fontFamilyShuffler,
		fontSizeSetter,
	)

	// default active component
	setThemeShufflerActive()

	for {
		e := <-eventChannel
		switch e {
		case "q", "<C-c>":
			return
		case "<Right>":
			setFontFamilyShufflerActive()
		case "<Left>":
			setThemeShufflerActive()
		case "<Tab>":
			setFontSetterActive()
		}
	}
}
