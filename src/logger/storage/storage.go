package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"time"
)

const folderName = ".idly"
const defaultContent = "[]"

var folderPath = homeFolder() + "/" + folderName

// TODO:
// □ file name template
// □ a file for every day? maybe to much? a month?
// □ create file if doesnot exists
// □

// CreateFolderIfNotExists _
func CreateFolderIfNotExists() {
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			log.Printf("createFolder error: %v", err)
		}
	}
}

func CreateFileIfNotExists() {
	CreateFolderIfNotExists()

	filePath := folderPath + "/" + fileName()
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		err = os.MkdirAll(folderPath, os.ModePerm)
		if err != nil {
			log.Printf("createFolder error: %v", err)
		}
	}
}

// Write _
func Write(selections []Selection) {
	content, _ := json.Marshal(selections)
	writeFile(filePath, content)
}

// Read _
func Read() []Selection {
	var content = readFile(filePath)
	var selections = make([]Selection, 0)
	err := json.Unmarshal(content, &selections)
	if err != nil {
		panic(err)
	}
	return selections
}

// Selection _
type Selection struct {
	Path     string   `json:"path"`
	Services []string `json:"services"`
}

func writeFile(path string, content []byte) {

	err := ioutil.WriteFile(path, []byte(content), 0644)
	if err != nil {
		panic(err)
	}
}

func readFile(path string) []byte {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return []byte(defaultContent)
	}
	return content
}

func homeFolder() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func fileName() string {
	now := time.Now()
	date := now.Format("2006-01-02")
	name := fmt.Sprintf("log-%s.json", date)

	return name
}
