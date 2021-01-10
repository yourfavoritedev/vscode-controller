package ui

import (
	ui "github.com/gizak/termui/v3"
)

// UIWidget implements methods to adapt the current ui.Drawable
type UIWidget interface {
	GetWidget() ui.Drawable
	InitWidget(*map[string]interface{})
	SetState() string
	ToggleActive()
}
