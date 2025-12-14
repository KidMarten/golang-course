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

func ReadFile(name string) ([]byte, error) {
	if isJSONFile(name) {
		data, err := os.ReadFile(name)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		return data, nil
	} else {
		return nil, fmt.Errorf("Not a json file")
	}
}
