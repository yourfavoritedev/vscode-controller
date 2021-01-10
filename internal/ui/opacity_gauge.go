package ui

import (
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	utils "github.com/yourfavoritedev/vscode-controller/utils"
)

const opacityProperty = "glassit.alpha"

// OpacityGauge is a structure for a widget that displays an opacity gauge
type OpacityGauge struct {
	Widget         *widgets.Gauge
	x1, y1, x2, y2 int
	settingsConfig *map[string]interface{}
}

func (og *OpacityGauge) toggleOpacity(key string) float64 {
	const maxOpacity = 255
	const lowestOpacity = 1
	const incrementer = 5
	currentOpacity := (*og.settingsConfig)[opacityProperty].(float64)
	if key == "a" {
		currentOpacity -= incrementer
	} else {
		currentOpacity += incrementer
	}

	if currentOpacity < lowestOpacity {
		currentOpacity = lowestOpacity
	} else if currentOpacity > maxOpacity {
		currentOpacity = maxOpacity
	}
	return currentOpacity
}

// GetWidget returns the underlying widget type
func (og *OpacityGauge) GetWidget() ui.Drawable {
	var widget ui.Drawable = og.Widget
	return widget
}

// ToggleActive changes the color of the widget to identify its status
func (og *OpacityGauge) ToggleActive() {
	if fgColor := og.Widget.BorderStyle.Fg; fgColor == ui.ColorYellow {
		og.Widget.BorderStyle.Fg = ui.ColorWhite
	} else {
		og.Widget.BorderStyle.Fg = ui.ColorYellow
	}
	ui.Render(og.Widget)
}

// SetState will navigate the current component
func (og *OpacityGauge) SetState() string {
	for {
		uiEvents := ui.PollEvents()
		e := <-uiEvents
		switch e.ID {
		case "<C-c>", "q", "A", "S", "W", "D":
			return e.ID
		case "a", "d":
			newOpacity := og.toggleOpacity(e.ID)
			newPercent := int(newOpacity / float64(255) * 100)
			og.Widget.Percent = newPercent

			(*og.settingsConfig)[opacityProperty] = newOpacity
			utils.WriteFileJSON(og.settingsConfig)

			ui.Render(og.Widget)
		}
	}
}

// InitWidget initializes the widget with preset settings
func (og *OpacityGauge) InitWidget(settings *map[string]interface{}) {
	og.settingsConfig = settings

	currentOpacity := (*settings)[opacityProperty].(float64)

	og.Widget.Title = "Opacity"
	og.Widget.Percent = int(currentOpacity / float64(255) * 100)
	og.Widget.BarColor = ui.ColorYellow
	og.Widget.SetRect(og.x1, og.y1, og.x2, og.y2)
	og.Widget.BorderStyle.Fg = ui.ColorWhite
	og.Widget.LabelStyle = ui.NewStyle(ui.ColorWhite)
}

// NewOpacityGauge creates an instance of OpacityGauge
func NewOpacityGauge(x1, y1, x2, y2 int) *OpacityGauge {
	return &OpacityGauge{
		Widget: widgets.NewGauge(),
		x1:     x1,
		y1:     y1,
		x2:     x2,
		y2:     y2,
	}
}
