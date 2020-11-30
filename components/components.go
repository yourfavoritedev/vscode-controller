package components

import (
	"fmt"
	"log"
	"os"
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/joho/godotenv"
	"github.com/yourfavoritedev/vscode-controller/fonts"
	"github.com/yourfavoritedev/vscode-controller/themes"
	utils "github.com/yourfavoritedev/vscode-controller/utils"
)

const (
	themeProperty      = "workbench.colorTheme"
	fontFamilyProperty = "editor.fontFamily"
	fontSizeProperty   = "editor.fontSize"
)

var settingsFilePath string
var jsonFile map[string]interface{}

func init() {
	// load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env files")
	}
	vsCodePath, ok := os.LookupEnv("VSCODE_PATH")

	if !ok {
		log.Fatalln("There was an error finding your vscode path")
	}

	settingsFilePath = fmt.Sprintf("%s/%s", vsCodePath, "settings.json")
	jsonFile = utils.GetFileJSON(settingsFilePath)
}

func changeProperty(propertyKey, selectedValue string) {
	jsonFile[propertyKey] = selectedValue
	utils.WriteFileJSON(settingsFilePath, &jsonFile)
}

func getSelectedRow(rows []string, value string) int {
	var selectedRow int
	for i, v := range rows {
		if v == value {
			selectedRow = i
		}
	}
	return selectedRow
}

// ThemeShuffler renders a list of themes for the user to select
func ThemeShuffler(eventChannel chan string) (*widgets.List, func() string) {
	var x1, y1, x2, y2 int = 0, 2, 30, 10
	// create theme list
	var rows = themes.AllThemes
	themesList := widgets.NewList()
	themesList.Title = "Themes"
	themesList.Rows = rows

	themesList.TextStyle = ui.NewStyle(ui.ColorYellow)
	themesList.WrapText = false

	selectedRow := getSelectedRow(rows, jsonFile[themeProperty].(string))
	themesList.SelectedRow = selectedRow
	themesList.SetRect(x1, y1, x2, y2)
	themesList.BorderStyle.Fg = ui.ColorWhite

	setActiveWidget := func() string {
		uiEvents := ui.PollEvents()

		for {
			// receive value from channel and set to e
			e := <-uiEvents
			switch e.ID {
			case "q", "<C-c>":
				return e.ID
			case "j", "<Down>":
				themesList.ScrollDown()
			case "k", "<Up>":
				themesList.ScrollUp()
			case "<Right>", "<Tab>":
				go func() { eventChannel <- e.ID }()
				return e.ID
			}

			selectedTheme := themesList.Rows[themesList.SelectedRow]
			changeProperty(themeProperty, selectedTheme)

			ui.Render(themesList)
		}
	}
	return themesList, setActiveWidget
}

// FontShuffler renders a list of fonts for the user to select
func FontShuffler(eventChannel chan string) (*widgets.List, func() string) {
	var x1, y1, x2, y2 int = 33, 2, 63, 10
	// create font list
	var rows = fonts.AllFonts
	fontsList := widgets.NewList()
	fontsList.Title = "Font Family"
	fontsList.Rows = rows

	fontsList.TextStyle = ui.NewStyle(ui.ColorYellow)
	fontsList.WrapText = false

	selectedRow := getSelectedRow(rows, jsonFile[fontFamilyProperty].(string))
	fontsList.SelectedRow = selectedRow
	fontsList.SetRect(x1, y1, x2, y2)
	fontsList.BorderStyle.Fg = ui.ColorWhite

	setActiveWidget := func() string {
		uiEvents := ui.PollEvents()

		for {
			// receive value from channel and set to e
			e := <-uiEvents
			switch e.ID {
			case "q", "<C-c>":
				return e.ID
			case "j", "<Down>":
				fontsList.ScrollDown()
			case "k", "<Up>":
				fontsList.ScrollUp()
			case "<Left>", "<Tab>":
				go func() { eventChannel <- e.ID }()
				return e.ID
			}

			selectedFont := fontsList.Rows[fontsList.SelectedRow]
			changeProperty(fontFamilyProperty, selectedFont)

			ui.Render(fontsList)
		}
	}
	return fontsList, setActiveWidget
}

// FontSizeSetter renders a list of sizes for the user to select
func FontSizeSetter(eventChannel chan string) (*widgets.List, func() string) {
	var x1, y1, x2, y2 int = 66, 2, 96, 10
	//create size list
	rows := make([]string, 29, 29)
	for i := 8; i < 37; i++ {
		s := strconv.Itoa(i)
		rows[i-8] = s
	}

	fontSize := widgets.NewList()
	fontSize.Title = "Font Size"
	fontSize.Rows = rows

	fontSize.TextStyle = ui.NewStyle(ui.ColorYellow)
	fontSize.WrapText = false

	selectedRow := getSelectedRow(rows, jsonFile[fontSizeProperty].(string))
	fontSize.SelectedRow = selectedRow
	fontSize.SetRect(x1, y1, x2, y2)
	fontSize.BorderStyle.Fg = ui.ColorWhite

	setActiveWidget := func() string {
		uiEvents := ui.PollEvents()

		for {
			// receive value from channel and set to e
			e := <-uiEvents
			switch e.ID {
			case "q", "<C-c>":
				return e.ID
			case "j", "<Down>":
				fontSize.ScrollDown()
			case "k", "<Up>":
				fontSize.ScrollUp()
			case "<Left>", "<Right>":
				go func() { eventChannel <- e.ID }()
				return e.ID
			}

			selectedFont := fontSize.Rows[fontSize.SelectedRow]
			changeProperty(fontSizeProperty, selectedFont)

			ui.Render(fontSize)
		}
	}

	return fontSize, setActiveWidget
}
