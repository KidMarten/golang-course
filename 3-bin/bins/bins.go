package bins

import (
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	bins []Bin
}

func NewBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	newBin := &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
	return newBin, nil
}

func NewBinList(bins []Bin) (*BinList, error) {
	newBinList := &BinList{
		bins: bins,
	}
	return newBinList, nil
}
