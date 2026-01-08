package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
			fmt.Println(err)
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
		fmt.Println(err)
	}
	_, err = file.Write(content)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Wirte successful")
}
