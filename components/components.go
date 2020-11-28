package components

import (
	"fmt"
	"log"
	"os"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/yourfavoritedev/terminal-app/themes"
	utils "github.com/yourfavoritedev/terminal-app/utils"
)

const themeProperty = "workbench.colorTheme"

func changeTheme(newTheme string) {
	vsCodePath, ok := os.LookupEnv("VSCODE_PATH")
	if !ok {
		log.Fatalln("There was an error finding your vscode path")
	}

	settingsFilePath := fmt.Sprintf("%s/%s", vsCodePath, "settings.json")
	jsonFile := utils.GetFileJSON(settingsFilePath)

	jsonFile[themeProperty] = newTheme

	utils.WriteFileJSON(settingsFilePath, jsonFile)
}

// ThemeShuffler renders a list of themes for the user to select
func ThemeShuffler() (*widgets.List, func() string) {
	// create theme list
	var rows []string
	for k := range themes.AllThemes {
		rows = append(rows, k)
	}

	themesList := widgets.NewList()
	themesList.Title = "Themes"
	themesList.Rows = rows

	themesList.TextStyle = ui.NewStyle(ui.ColorYellow)
	themesList.WrapText = false
	themesList.SetRect(0, 3, 35, 10)
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
			}

			selectedTheme := themesList.Rows[themesList.SelectedRow]
			changeTheme(selectedTheme)

			ui.Render(themesList)
		}
	}
	return themesList, setActiveWidget
}
