package ui

import (
	"sort"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/yourfavoritedev/vscode-controller/internal/fonts"
	utils "github.com/yourfavoritedev/vscode-controller/utils"
)

const fontFamilyProperty = "editor.fontFamily"

// FontFamilyShuffler is a structure for a widget that displays a list of font-familys
type FontFamilyShuffler struct {
	Widget         *widgets.List
	x1, y1, x2, y2 int
	settingsConfig *map[string]interface{}
}

// GetWidget returns the underlying widget type
func (ff *FontFamilyShuffler) GetWidget() ui.Drawable {
	var widget ui.Drawable = ff.Widget
	return widget
}

// ToggleActive changes the color of the widget to identify its status
func (ff *FontFamilyShuffler) ToggleActive() {
	if fgColor := ff.Widget.BorderStyle.Fg; fgColor == ui.ColorYellow {
		ff.Widget.BorderStyle.Fg = ui.ColorWhite
	} else {
		ff.Widget.BorderStyle.Fg = ui.ColorYellow
	}
	ui.Render(ff.Widget)
}

// InitWidget initializes the widget with preset settings
func (ff *FontFamilyShuffler) InitWidget(settings *map[string]interface{}) {
	ff.settingsConfig = settings

	rows := make([]string, len(fonts.AllFonts))
	for i, theme := range fonts.AllFonts {
		rows[i] = theme
	}
	sort.Strings(rows)

	var selectedRow int
	currentFontFamily := (*settings)[fontFamilyProperty].(string)
	for i, theme := range rows {
		if theme == currentFontFamily {
			selectedRow = i
		}
	}

	ff.Widget.Title = "Font Family"
	ff.Widget.Rows = rows
	ff.Widget.SelectedRow = selectedRow
	ff.Widget.TextStyle = ui.NewStyle(ui.ColorWhite)
	ff.Widget.WrapText = false
	ff.Widget.SelectedRowStyle.Fg = ui.ColorYellow
	ff.Widget.SetRect(ff.x1, ff.y1, ff.x2, ff.y2)
	ff.Widget.BorderStyle.Fg = ui.ColorWhite
}

// SetState will navigate the current component
func (ff *FontFamilyShuffler) SetState() string {
	for {
		uiEvents := ui.PollEvents()
		e := <-uiEvents
		switch e.ID {
		case "<C-c>", "q", "A", "S", "W", "D":
			return e.ID
		case "s":
			ff.Widget.ScrollDown()
		case "w":
			ff.Widget.ScrollUp()
		}

		selectedFontFamily := ff.Widget.Rows[ff.Widget.SelectedRow]
		(*ff.settingsConfig)[fontFamilyProperty] = selectedFontFamily
		utils.WriteFileJSON(ff.settingsConfig)

		ui.Render(ff.Widget)
	}
}

// NewFontFamilyShuffler creates an instance of FontFamilyShuffler
func NewFontFamilyShuffler(x1, y1, x2, y2 int) *FontFamilyShuffler {
	return &FontFamilyShuffler{
		Widget: widgets.NewList(),
		x1:     x1,
		y1:     y1,
		x2:     x2,
		y2:     y2,
	}
}
