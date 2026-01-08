package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
)

type FileStorage struct {
	filename string
}

func NewFileStorage(filename string) *FileStorage {
	return &FileStorage{filename: filename}
}

func (storage *FileStorage) Read() ([]byte, error) {
	if isJSONFile(storage.filename) {
		data, err := os.ReadFile(storage.filename)
		if err != nil {
			color.Yellow(err.Error())
			return nil, err
		}
		return data, nil
	} else {
		return nil, fmt.Errorf("Not a json file")
	}
}

func isJSONFile(filename string) bool {
	ext := filepath.Ext(filename)
	return strings.ToLower(ext) == ".json"
}

func (storage *FileStorage) Write(content []byte) {
	file, err := os.Create(storage.filename)
	if err != nil {
		color.Red("Unable to create file %s", err.Error())
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		color.Red("Unable to write file %s", err.Error())
		return
	}
	color.Green("Wirte successful")
}
