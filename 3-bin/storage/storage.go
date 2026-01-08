package storage

import (
	"app/bin/bins"
	"encoding/json"

	"github.com/fatih/color"
)

type DbProvider interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Storage struct {
	DbProvider DbProvider
	BinList    bins.BinList
}

func NewStorage(dbProvider DbProvider) *Storage {
	file, err := dbProvider.Read()
	if err != nil {
		color.Yellow("No data exists. Creating an empty storage.")
		return &Storage{
			DbProvider: dbProvider,
		}
	}
	var binList bins.BinList
	err = json.Unmarshal(file, &binList)
	if err != nil {
		color.Red("Failed to read data from storage. Returning an empty storage.")
		return &Storage{
			DbProvider: dbProvider,
		}
	}
	return &Storage{
		DbProvider: dbProvider,
		BinList:    binList,
	}
}

func (storage *Storage) ToBytes() ([]byte, error) {
	file, err := json.Marshal(storage.BinList)
	if err != nil {
		color.Red(err.Error())
		return nil, err
	}
	return file, nil
}

// Сохранение bin в виде json в локальном файле
func (storage *Storage) saveBinList() {
	data, err := storage.ToBytes()
	if err != nil {
		color.Red("Unable to save bin list. Failed to convert json to bytes")
	} else {
		storage.DbProvider.Write(data)
	}
}

func (storage *Storage) readBinList() (*bins.BinList, error) {
	file, err := storage.DbProvider.Read()
	var binList bins.BinList
	err = json.Unmarshal(file, &binList)
	if err != nil {
		return nil, err
	}
	return &binList, nil
}
