package assignment

import (
	"stock.ngmengineering.fr/stock/item"
)

type Unit string
type Quantity uint

type Assignment struct {
	Item     item.Item
	Unit     Unit
	Quantity Quantity
	Uptake   []Uptake
}

func NewAssignment(
	item item.Item,
	unit Unit,
	quantity Quantity,
) Assignment {
	return Assignment{
		Item:     item,
		Unit:     unit,
		Quantity: quantity,
	}
}

func (assignment *Assignment) GetAvailableQuantity() Quantity {
	availableQuantity := assignment.Quantity

	for _, uptake := range assignment.Uptake {
		availableQuantity -= uptake.Quantity
	}

	return availableQuantity
}
