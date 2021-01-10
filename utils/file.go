package file

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// GetFileJSON returns a populated JSONFile from filepath
func GetFileJSON(filepath string) map[string]interface{} {
	// open json file
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// read our initialized file as a byte array
	byteValue, _ := ioutil.ReadAll(file)

	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)
	return result
}

// WriteFileJSON converts newContent into json to be written to the filepath
func WriteFileJSON(newContent *map[string]interface{}) {
	var settingsFilePath string
	vsCodePath, ok := os.LookupEnv("VSCODE_PATH")

	if !ok {
		log.Fatalln("There was an error finding your vscode path")
	}

	settingsFilePath = fmt.Sprintf("%s/%s", vsCodePath, "settings.json")

	newJSON, err := json.Marshal(newContent)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(settingsFilePath, newJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
