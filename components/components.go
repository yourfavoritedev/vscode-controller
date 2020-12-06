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
func ThemeShuffler() (*widgets.List, func() string) {
	var x1, y1, x2, y2 int = 0, 2, 30, 10
	// create theme list
	var rows = themes.AllThemes
	themesList := widgets.NewList()
	themesList.Title = "Themes"
	themesList.Rows = rows

	themesList.TextStyle = ui.NewStyle(ui.ColorWhite)
	themesList.WrapText = false

	selectedRow := getSelectedRow(rows, jsonFile[themeProperty].(string))
	themesList.SelectedRow = selectedRow
	themesList.SelectedRowStyle.Fg = ui.ColorYellow
	themesList.SetRect(x1, y1, x2, y2)
	themesList.BorderStyle.Fg = ui.ColorWhite

	setActiveComponent := func() string {
		uiEvents := ui.PollEvents()

		for {
			e := <-uiEvents
			switch e.ID {
			case "<C-c>", "A", "S", "W", "D":
				return e.ID
			case "s":
				themesList.ScrollDown()
			case "w":
				themesList.ScrollUp()
			}

			selectedTheme := themesList.Rows[themesList.SelectedRow]
			changeProperty(themeProperty, selectedTheme)

			ui.Render(themesList)
		}
	}
	return themesList, setActiveComponent
}

// FontShuffler renders a list of fonts for the user to select
func FontShuffler() (*widgets.List, func() string) {
	var x1, y1, x2, y2 int = 33, 2, 63, 10
	// create font list
	var rows = fonts.AllFonts
	fontsList := widgets.NewList()
	fontsList.Title = "Font Family"
	fontsList.Rows = rows

	fontsList.TextStyle = ui.NewStyle(ui.ColorWhite)
	fontsList.WrapText = false

	selectedRow := getSelectedRow(rows, jsonFile[fontFamilyProperty].(string))
	fontsList.SelectedRow = selectedRow
	fontsList.SelectedRowStyle.Fg = ui.ColorYellow
	fontsList.SetRect(x1, y1, x2, y2)
	fontsList.BorderStyle.Fg = ui.ColorWhite

	setActiveComponent := func() string {
		uiEvents := ui.PollEvents()

		for {
			e := <-uiEvents

			switch e.ID {
			case "<C-c>", "A", "S", "W", "D":
				return e.ID
			case "s":
				fontsList.ScrollDown()
			case "w":
				fontsList.ScrollUp()
			}

			selectedFont := fontsList.Rows[fontsList.SelectedRow]
			changeProperty(fontFamilyProperty, selectedFont)

			ui.Render(fontsList)
		}
	}
	return fontsList, setActiveComponent
}

// FontSizeSetter renders a list of sizes for the user to select
func FontSizeSetter() (*widgets.List, func() string) {
	var x1, y1, x2, y2 int = 0, 11, 30, 21
	//create size list
	rows := make([]string, 29, 29)
	for i := 8; i < 37; i++ {
		s := strconv.Itoa(i)
		rows[i-8] = s
	}

	fontSize := widgets.NewList()
	fontSize.Title = "Font Size"
	fontSize.Rows = rows

	fontSize.TextStyle = ui.NewStyle(ui.ColorWhite)
	fontSize.WrapText = false

	selectedRow := getSelectedRow(rows, jsonFile[fontSizeProperty].(string))
	fontSize.SelectedRow = selectedRow
	fontSize.SelectedRowStyle.Fg = ui.ColorYellow
	fontSize.SetRect(x1, y1, x2, y2)
	fontSize.BorderStyle.Fg = ui.ColorWhite

	setActiveComponent := func() string {
		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "<C-c>", "A", "S", "W", "D":
				return e.ID
			case "s":
				fontSize.ScrollDown()
			case "w":
				fontSize.ScrollUp()
			}

			selectedFont := fontSize.Rows[fontSize.SelectedRow]
			changeProperty(fontSizeProperty, selectedFont)

			ui.Render(fontSize)
		}
	}

	return fontSize, setActiveComponent
}

// TempListSetter renders a list of sizes for the user to select
func TempListSetter() (*widgets.List, func() string) {
	var x1, y1, x2, y2 int = 33, 11, 63, 21
	//create size list
	rows := make([]string, 29, 29)
	for i := 8; i < 37; i++ {
		s := strconv.Itoa(i)
		rows[i-8] = s
	}

	fontSize := widgets.NewList()
	fontSize.Title = "Temp List"
	fontSize.Rows = rows

	fontSize.TextStyle = ui.NewStyle(ui.ColorWhite)
	fontSize.WrapText = false
	selectedRow := getSelectedRow(rows, jsonFile[fontSizeProperty].(string))
	fontSize.SelectedRow = selectedRow
	fontSize.SelectedRowStyle.Fg = ui.ColorYellow

	fontSize.SetRect(x1, y1, x2, y2)
	fontSize.BorderStyle.Fg = ui.ColorWhite

	setActiveComponent := func() string {
		uiEvents := ui.PollEvents()
		for {
			e := <-uiEvents
			switch e.ID {
			case "<C-c>", "A", "S", "W", "D":
				return e.ID
			case "s":
				fontSize.ScrollDown()
			case "w":
				fontSize.ScrollUp()
			}

			selectedFont := fontSize.Rows[fontSize.SelectedRow]
			changeProperty(fontSizeProperty, selectedFont)

			ui.Render(fontSize)
		}
	}

	return fontSize, setActiveComponent
}

// ThemeShuffler2 renders a list of themes for the user to select
func ThemeShuffler2() (*widgets.List, func() string) {
	var x1, y1, x2, y2 int = 66, 2, 96, 10
	// create theme list
	var rows = themes.AllThemes
	themesList := widgets.NewList()
	themesList.Title = "Themes2"
	themesList.Rows = rows

	themesList.TextStyle = ui.NewStyle(ui.ColorWhite)
	themesList.WrapText = false

	selectedRow := getSelectedRow(rows, jsonFile[themeProperty].(string))
	themesList.SelectedRow = selectedRow
	themesList.SelectedRowStyle.Fg = ui.ColorYellow
	themesList.SetRect(x1, y1, x2, y2)
	themesList.BorderStyle.Fg = ui.ColorWhite

	setActiveComponent := func() string {
		uiEvents := ui.PollEvents()

		for {
			e := <-uiEvents
			switch e.ID {
			case "<C-c>", "A", "S", "W", "D":
				return e.ID
			case "s":
				themesList.ScrollDown()
			case "w":
				themesList.ScrollUp()
			}

			selectedTheme := themesList.Rows[themesList.SelectedRow]
			changeProperty(themeProperty, selectedTheme)

			ui.Render(themesList)
		}
	}
	return themesList, setActiveComponent
}

// OpacityGauge renders a toggle for the user to adjust their opacity
func OpacityGauge() (*widgets.Gauge, func() string) {
	var x1, y1, x2, y2 int = 0, 23, 75, 27
	gauge := widgets.NewGauge()
	gauge.Title = "Opacity"
	gauge.SetRect(x1, y1, x2, y2)
	gauge.Percent = 75
	gauge.BarColor = ui.ColorYellow
	gauge.BorderStyle.Fg = ui.ColorWhite
	gauge.LabelStyle = ui.NewStyle(ui.ColorWhite)

	setActiveComponent := func() string {
		uiEvents := ui.PollEvents()

		for {
			e := <-uiEvents
			switch e.ID {
			case "<C-c>", "A", "S", "W", "D":
				return e.ID
			case "a":
				// reduce opacity
			case "d":
				// increase opacity
			}
			ui.Render(gauge)
		}
	}

	return gauge, setActiveComponent
}
