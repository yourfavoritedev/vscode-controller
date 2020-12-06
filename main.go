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
	TempListSetter, TempListSetterActive := components.TempListSetter()
	ThemeShuffler2, setThemeShufflerActive2 := components.ThemeShuffler2()
	opacityGauge, setOpacityGaugeActive := components.OpacityGauge()

	// set default active component
	var activeComponent ui.Drawable = themeShuffler
	activeComponent.(*widgets.List).BorderStyle.Fg = ui.ColorYellow
	activeComponent.(*widgets.List).TitleStyle.Fg = ui.ColorYellow

	// render the initial ui
	ui.Render(
		themeShuffler,
		fontFamilyShuffler,
		ThemeShuffler2,
		fontSizeSetter,
		TempListSetter,
		opacityGauge,
	)

	rowOne := []ui.Drawable{
		themeShuffler,
		fontFamilyShuffler,
		ThemeShuffler2,
	}

	rowTwo := []ui.Drawable{
		fontSizeSetter,
		TempListSetter,
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

		// focus activeComponent via type switch
		switch ac := activeComponent.(type) {
		case *widgets.List:
			ac.BorderStyle.Fg = ui.ColorYellow
			ac.TitleStyle.Fg = ui.ColorYellow
			ui.Render(ac)
			activeComponent = ac

			switch ac.Title {
			case "Font Family":
				e = setFontShufflerActive()
			case "Themes":
				e = setThemeShufflerActive()
			case "Font Size":
				e = setFontSetterActive()
			case "Temp List":
				e = TempListSetterActive()
			case "Themes2":
				e = setThemeShufflerActive2()
			}
		case *widgets.Gauge:
			ac.BorderStyle.Fg = ui.ColorYellow
			ac.TitleStyle.Fg = ui.ColorYellow
			ui.Render(ac)
			activeComponent = ac
			switch ac.Title {
			case "Opacity Gauage":
				e = setOpacityGaugeActive()
			}
		}
	}
}
