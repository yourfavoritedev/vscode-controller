package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/yourfavoritedev/vscode-controller/internal/ui"
	utils "github.com/yourfavoritedev/vscode-controller/utils"
)

var settingsFilePath string
var settingsFile map[string]interface{}

func init() {
	//load environment variables
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env files")
	}
	vsCodePath, ok := os.LookupEnv("VSCODE_PATH")

	if !ok {
		log.Fatalln("There was an error finding your vscode path")
	}

	settingsFilePath = fmt.Sprintf("%s/%s", vsCodePath, "settings.json")
	settingsFile = utils.GetFileJSON(settingsFilePath)
}

func main() {
	row1 := []ui.UIWidget{
		ui.NewThemeShuffler(0, 0, 25, 10),
		ui.NewFontFamilyShuffler(26, 0, 51, 10),
	}

	row2 := []ui.UIWidget{
		ui.NewFontSizeSetter(0, 11, 25, 21),
	}

	row3 := []ui.UIWidget{
		ui.NewOpacityGauge(0, 22, 75, 25),
	}

	ui.WidgetsController(
		&settingsFile,
		row1,
		row2,
		row3,
	)
}
