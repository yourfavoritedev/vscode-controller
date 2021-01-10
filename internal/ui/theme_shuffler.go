package ui

import (
	"sort"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/yourfavoritedev/vscode-controller/internal/themes"
	utils "github.com/yourfavoritedev/vscode-controller/utils"
)

const themeProperty = "workbench.colorTheme"

// ThemeShuffler is a structure for a widget that displays a list of themes
type ThemeShuffler struct {
	Widget         *widgets.List
	x1, y1, x2, y2 int
	settingsConfig *map[string]interface{}
}

// GetWidget returns the underlying widget type
func (ts *ThemeShuffler) GetWidget() ui.Drawable {
	var widget ui.Drawable = ts.Widget
	return widget
}

// ToggleActive changes the color of the widget to identify its status
func (ts *ThemeShuffler) ToggleActive() {
	if fgColor := ts.Widget.BorderStyle.Fg; fgColor == ui.ColorYellow {
		ts.Widget.BorderStyle.Fg = ui.ColorWhite
	} else {
		ts.Widget.BorderStyle.Fg = ui.ColorYellow
	}
	ui.Render(ts.Widget)
}

// InitWidget initializes the widget with preset settings
func (ts *ThemeShuffler) InitWidget(settings *map[string]interface{}) {
	ts.settingsConfig = settings

	rows := make([]string, len(themes.AllThemes))
	for i, theme := range themes.AllThemes {
		rows[i] = theme
	}
	sort.Strings(rows)

	var selectedRow int
	currentTheme := (*settings)[themeProperty].(string)
	for i, theme := range rows {
		if theme == currentTheme {
			selectedRow = i
		}
	}

	ts.Widget.Title = "Themes"
	ts.Widget.Rows = rows
	ts.Widget.SelectedRow = selectedRow
	ts.Widget.TextStyle = ui.NewStyle(ui.ColorWhite)
	ts.Widget.WrapText = false
	ts.Widget.SelectedRowStyle.Fg = ui.ColorYellow
	ts.Widget.SetRect(ts.x1, ts.y1, ts.x2, ts.y2)
	ts.Widget.BorderStyle.Fg = ui.ColorWhite
}

// SetState will navigate the current component
func (ts *ThemeShuffler) SetState() string {
	for {
		uiEvents := ui.PollEvents()
		e := <-uiEvents
		switch e.ID {
		case "<C-c>", "q", "A", "S", "W", "D":
			return e.ID
		case "s":
			ts.Widget.ScrollDown()
		case "w":
			ts.Widget.ScrollUp()
		}

		selectedTheme := ts.Widget.Rows[ts.Widget.SelectedRow]
		(*ts.settingsConfig)[themeProperty] = selectedTheme
		utils.WriteFileJSON(ts.settingsConfig)

		ui.Render(ts.Widget)
	}
}

// NewThemeShuffler creates an instance of ThemeShuffler
func NewThemeShuffler(x1, y1, x2, y2 int) *ThemeShuffler {
	return &ThemeShuffler{
		Widget: widgets.NewList(),
		x1:     x1,
		y1:     y1,
		x2:     x2,
		y2:     y2,
	}
}
