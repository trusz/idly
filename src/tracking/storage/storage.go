package storage

import (
	"io/ioutil"
)

// Storage _
type Storage struct {
	storageFilePath string
}

// New _
func New() Storage {
	var fileName = dateStampedFileName()
	ensureFileExists(fileName, folderPath)
	return Storage{
		storageFilePath: filePath(fileName, folderPath),
	}
}

func (s Storage) Write(p []byte) (n int, err error) {
	err = ioutil.WriteFile(s.storageFilePath, p, 0644)
	if err != nil {
		return 0, err
	}
	return len(p), nil
}
