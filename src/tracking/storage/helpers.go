package storage

import (
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

func ensureFolderExists(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Printf("createFolder error: %v", err)
		}
	}
}

func ensureFileExists(fileName string, inFolder string) {
	ensureFolderExists(inFolder)

	filePath := filePath(fileName, inFolder)

	if !fileExists(filePath) {
		createFile(filePath)
	}

}

func fileExists(filePath string) bool {

	var _, err = os.Stat(filePath)
	return !os.IsNotExist(err)
}

func createFile(filePath string) {
	var file, err = os.Create(filePath)
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()
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

func dateStampedFileName() string {
	now := time.Now()
	date := now.Format("2006-01")
	name := fmt.Sprintf("log-%s.txt", date)

	return name
}

func filePath(fileName string, folderPath string) string {
	return folderPath + "/" + fileName
}
