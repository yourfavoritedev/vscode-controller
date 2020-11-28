package file

import (
	"encoding/json"
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
func WriteFileJSON(filepath string, newContent interface{}) {
	newJSON, err := json.Marshal(newContent)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(filepath, newJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
