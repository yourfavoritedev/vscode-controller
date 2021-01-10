package ui

import (
	"log"
	"math"

	ui "github.com/gizak/termui/v3"
)

func init() {
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
}

// WidgetsController controls the navigation and activation of all widgets
func WidgetsController(settings *map[string]interface{}, rows ...[]UIWidget) {
	defer ui.Close()

	numWidgets := 0
	for _, row := range rows {
		numWidgets += len(row)
	}

	widgets := make([]UIWidget, numWidgets)
	drawables := make([]ui.Drawable, numWidgets)

	// initialize all components
	i := 0
	for _, row := range rows {
		for _, widget := range row {
			widget.InitWidget(settings)
			drawables[i] = widget.GetWidget()
			widgets[i] = widget
			i++
		}
	}

	ui.Render(drawables...)

	// First widget -> default widget
	activeWidget := widgets[0]
	activeWidget.ToggleActive()

	activeRowIndex, activeColumnIndex := 0, 0

	e := activeWidget.SetState()
	for {
		switch e {
		case "A":
			activeColumnIndex--
		case "S":
			activeRowIndex++
		case "D":
			activeColumnIndex++
		case "W":
			activeRowIndex--
		case "<C-c>", "q":
			return
		}

		// Deactive current active widget
		activeWidget.ToggleActive()

		// reposition if cursor is out of range
		if activeRowIndex < 0 {
			absActiveRowIndex := int(math.Abs(float64(activeRowIndex)))
			activeRowIndex = len(rows) - absActiveRowIndex
		} else if activeRowIndex > len(rows)-1 {
			activeRowIndex = 0
		}
		activeRow := rows[activeRowIndex]

		if activeColumnIndex < 0 {
			absActiveColumnIndex := int(math.Abs(float64(activeColumnIndex)))
			activeColumnIndex = len(activeRow) - absActiveColumnIndex
		} else if activeColumnIndex >= len(activeRow) {
			activeColumnIndex = 0
		}
		activeColumn := activeRow[activeColumnIndex]

		// Set new active widget
		activeWidget = activeColumn
		activeWidget.ToggleActive()

		e = activeWidget.SetState()
	}
}
