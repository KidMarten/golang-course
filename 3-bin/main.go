package main

import (
	"app/bin/file"
	"app/bin/storage"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Yellow("__Bin Handler__")
	jsonStorage := file.NewFileStorage("data.json")
	storage := storage.NewStorage(jsonStorage)
	fmt.Println(storage.BinList)
}
