package utils

import (
	"errors"
	"log"
	"os"
)

func NewDirIfNotExists(dir string) string {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return ""
	} else {
		return dir
	}
}

func FileExists(name string) bool {
	_, err := os.Stat(name)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func WriteStrToFile(fileName string, s string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	} else {

		_, errFileWriteString := file.WriteString(s)
		if errFileWriteString != nil {
			file.Close()
			return
		}
	}
	file.Close()
}
