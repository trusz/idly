package storage

import (
	"os"
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

	f, err := os.OpenFile(s.storageFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return -1, nil
	}

	if _, err := f.Write(p); err != nil {
		return -1, nil
	}
	if err := f.Close(); err != nil {
		return -1, nil
	}

	return len(p), nil

}
