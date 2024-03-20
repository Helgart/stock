package store

import (
	"github.com/google/uuid"
	"stock.ngmengineering.fr/stock/assignment"
	"stock.ngmengineering.fr/stock/item"
)

type Store struct {
	Name        string
	Assignments map[uuid.UUID][]assignment.Assignment
}

func NewStore(name string) Store {
	return Store{
		Name:        name,
		Assignments: map[uuid.UUID][]assignment.Assignment{},
	}
}

func (store *Store) assignItem(
	item item.Item,
	unit assignment.Unit,
	quantity assignment.Quantity,
) assignment.Assignment {
	assignedItem := assignment.NewAssignment(item, unit, quantity)

	store.Assignments[item.Uid] = append(store.Assignments[item.Uid], assignedItem)

	return assignedItem
}
