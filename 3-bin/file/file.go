package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func isJSONFile(filename string) bool {
	ext := filepath.Ext(filename)
	return strings.ToLower(ext) == ".json"
}

func ReadFile(filename string) ([]byte, error) {
	if isJSONFile(filename) {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return data, nil
	} else {
		return nil, fmt.Errorf("Not a json file")
	}
}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
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
