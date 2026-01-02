package storage

import (
	"app/bin/bins"
	"app/bin/file"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
)

type Storage struct {
	BinList bins.BinList `json:"bin_list"`
}

func (storage *Storage) ToBytes() ([]byte, error) {
	file, err := json.Marshal(storage)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return file, nil
}

// Сохранение bin в виде json в локальном файле
func (storage *Storage) saveBinList(filename string) {
	data, err := storage.ToBytes()
	if err != nil {
		color.Red("Failed to convert json to bytes")
	}
	file.WriteFile(data, filename)
}

func (storage *Storage) readBinList(filename string) (*bins.BinList, error) {
	file, err := file.ReadFile(filename)
	var binList bins.BinList
	err = json.Unmarshal(file, &binList)
	if err != nil {
		return nil, err
	}
	return &binList, nil
}
