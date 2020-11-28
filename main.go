package main

import (
	"log"

	ui "github.com/gizak/termui/v3"
	"github.com/joho/godotenv"
	components "github.com/yourfavoritedev/terminal-app/components"
)

func init() {
	// load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env files")
	}
	// initializes ui
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
}

func main() {
	defer ui.Close()
	themeShuffler, setThemeShufflerActive := components.ThemeShuffler()
	ui.Render(themeShuffler)
	setThemeShufflerActive()
}
