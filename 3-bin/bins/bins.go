package bins

import (
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
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
		Bins: bins,
	}
	return newBinList, nil
}
