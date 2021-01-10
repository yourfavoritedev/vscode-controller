package ui

import (
	"strconv"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	utils "github.com/yourfavoritedev/vscode-controller/utils"
)

const fontSizeProperty = "editor.fontSize"

// FontSizeSetter is a structure for a widget that displays a list of font-sizes
type FontSizeSetter struct {
	Widget         *widgets.List
	x1, y1, x2, y2 int
	settingsConfig *map[string]interface{}
}

// GetWidget returns the underlying widget type
func (fs *FontSizeSetter) GetWidget() ui.Drawable {
	var widget ui.Drawable = fs.Widget
	return widget
}

// ToggleActive changes the color of the widget to identify its status
func (fs *FontSizeSetter) ToggleActive() {
	if fgColor := fs.Widget.BorderStyle.Fg; fgColor == ui.ColorYellow {
		fs.Widget.BorderStyle.Fg = ui.ColorWhite
	} else {
		fs.Widget.BorderStyle.Fg = ui.ColorYellow
	}
	ui.Render(fs.Widget)
}

// InitWidget initializes the widget with preset settings
func (fs *FontSizeSetter) InitWidget(settings *map[string]interface{}) {
	fs.settingsConfig = settings

	rows := make([]string, 30)
	for i, v := 0, 8; i < len(rows); i++ {
		s := strconv.Itoa(v)
		rows[i] = s
		v++
	}

	var selectedRow int
	currentFontSize := (*settings)[fontSizeProperty].(string)
	for i, size := range rows {
		if size == currentFontSize {
			selectedRow = i
		}
	}

	fs.Widget.Title = "Font Size"
	fs.Widget.Rows = rows
	fs.Widget.SelectedRow = selectedRow
	fs.Widget.TextStyle = ui.NewStyle(ui.ColorWhite)
	fs.Widget.WrapText = false
	fs.Widget.SelectedRowStyle.Fg = ui.ColorYellow
	fs.Widget.SetRect(fs.x1, fs.y1, fs.x2, fs.y2)
	fs.Widget.BorderStyle.Fg = ui.ColorWhite
}

// SetState will navigate the current component
func (fs *FontSizeSetter) SetState() string {
	for {
		uiEvents := ui.PollEvents()
		e := <-uiEvents
		switch e.ID {
		case "<C-c>", "q", "A", "S", "W", "D":
			return e.ID
		case "s":
			fs.Widget.ScrollDown()
		case "w":
			fs.Widget.ScrollUp()
		}

		selectedSize := fs.Widget.Rows[fs.Widget.SelectedRow]
		(*fs.settingsConfig)[fontSizeProperty] = selectedSize
		utils.WriteFileJSON(fs.settingsConfig)

		ui.Render(fs.Widget)
	}
}

// NewFontSizeSetter creates an instance of FontSizeSetter
func NewFontSizeSetter(x1, y1, x2, y2 int) *FontSizeSetter {
	return &FontSizeSetter{
		Widget: widgets.NewList(),
		x1:     x1,
		y1:     y1,
		x2:     x2,
		y2:     y2,
	}
}
