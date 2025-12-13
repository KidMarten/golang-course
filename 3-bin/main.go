package main

import (
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func newBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	newBin := &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
	return newBin, nil
}

func newBinList(bins []Bin) (*BinList, error) {
	newBinList := &BinList{
		bins: bins,
	}
	return newBinList, nil
}

func main() {

}
