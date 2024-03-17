package assignment

import (
	"stock.ngmengineering.fr/stock/item"
	"stock.ngmengineering.fr/stock/store"
)

type Unit string
type Quantity uint

type Assignment struct {
	Item     item.Item
	Store    store.Store
	Unit     Unit
	Quantity Quantity
}

func NewAssignment(
	item item.Item,
	store store.Store,
	unit Unit,
	quantity Quantity,
) Assignment {
	return Assignment{
		Item:     item,
		Store:    store,
		Unit:     unit,
		Quantity: quantity,
	}
}
