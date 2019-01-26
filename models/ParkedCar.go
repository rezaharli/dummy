package models

import (
	"fmt"
)

type ParkedCar struct {
	Color     string
	RegNumber string
	Slot      int
}

func (p *ParkedCar) String() string {
	return fmt.Sprintf("%d\t%s\t%s", p.Slot, p.RegNumber, p.Color)
}
