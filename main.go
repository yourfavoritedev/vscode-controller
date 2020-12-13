package main

import (
	"log"
	"math"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
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

	themeShuffler, setThemeShufflerActive := components.ThemeShuffler()
	fontFamilyShuffler, setFontShufflerActive := components.FontShuffler()
	fontSizeSetter, setFontSetterActive := components.FontSizeSetter()
	opacityGauge, setOpacityGaugeActive := components.OpacityGauge()

	// set default current component
	var currentComponent ui.Drawable = themeShuffler
	currentComponent.(*widgets.List).BorderStyle.Fg = ui.ColorYellow
	currentComponent.(*widgets.List).TitleStyle.Fg = ui.ColorYellow

	// render the initial ui
	ui.Render(
		themeShuffler,
		fontFamilyShuffler,
		fontSizeSetter,
		opacityGauge,
	)

	rowOne := []ui.Drawable{
		themeShuffler,
		fontFamilyShuffler,
	}

	rowTwo := []ui.Drawable{
		fontSizeSetter,
	}

	rowThree := []ui.Drawable{
		opacityGauge,
	}

	componentGrid := [][]ui.Drawable{
		rowOne,
		rowTwo,
		rowThree,
	}

	// set initial position
	var activeComponent ui.Drawable
	var activeRowIndex, activeColumnIndex int

	e := setThemeShufflerActive()

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
		case "<C-c>":
			return
		}

		// reposition if cursor is out of range
		if activeRowIndex < 0 {
			absRowIndex := int(math.Abs(float64(activeRowIndex)))
			activeRowIndex = len(componentGrid) - absRowIndex
		} else if activeRowIndex > len(componentGrid)-1 {
			activeRowIndex = 0
		}

		if activeColumnIndex < 0 {
			absColumnIndex := int(math.Abs(float64(activeColumnIndex)))
			activeColumnIndex = len(componentGrid[activeRowIndex]) - absColumnIndex
		} else if activeColumnIndex > len(componentGrid[activeRowIndex])-1 {
			activeColumnIndex = 0
		}

		activeComponent = componentGrid[activeRowIndex][activeColumnIndex]

		// deactivate current component
		switch cc := currentComponent.(type) {
		case *widgets.List:
			cc.BorderStyle.Fg = ui.ColorWhite
			cc.TitleStyle.Fg = ui.ColorWhite
			ui.Render(cc)
		case *widgets.Gauge:
			cc.BorderStyle.Fg = ui.ColorWhite
			cc.TitleStyle.Fg = ui.ColorWhite
			ui.Render(cc)
		}

		// focus active component via type switch
		switch ac := activeComponent.(type) {
		case *widgets.List:
			ac.BorderStyle.Fg = ui.ColorYellow
			ac.TitleStyle.Fg = ui.ColorYellow
			ui.Render(ac)
			currentComponent = ac

			switch ac.Title {
			case "Font Family":
				e = setFontShufflerActive()
			case "Themes":
				e = setThemeShufflerActive()
			case "Font Size":
				e = setFontSetterActive()
			}
		case *widgets.Gauge:
			ac.BorderStyle.Fg = ui.ColorYellow
			ac.TitleStyle.Fg = ui.ColorYellow
			ui.Render(ac)
			currentComponent = ac

			switch ac.Title {
			case "Opacity":
				e = setOpacityGaugeActive()
			}
		}
	}
}
