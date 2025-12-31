package storage

import (
	"app/bin/bins"
)

type Storage struct {
	Bins []bins.Bin `json:"bins"`
}
