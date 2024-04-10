package store

import (
	"github.com/Helgart/stock/assignment"
	"github.com/Helgart/stock/item"
	"github.com/google/uuid"
)

// Store is a type representing a store that holds assignments of items.
// Each store has a name and a map of assignments.
// The assignments represent the availability of an item in the store,
// including the unit, quantity, and uptakes.
//
// Example usage:
//
//	newStore := NewStore("MyStore")
//	newItem := item.Item{Uid: uuid.New(), Name: "Item1"}
//	newUnit := assignment.Unit("Piece")
//	newQuantity := assignment.Quantity(10)
//
//	newStore.AssignItem(newItem, newUnit, newQuantity)
//	availableQuantity := newStore.GetAvailableQuantityForItem(newItem)
type Store struct {
	Name        string
	Assignments map[uuid.UUID][]*assignment.Assignment
}

// NewStore initializes a new Store with the given name and an empty assignments map.
func NewStore(name string) *Store {
	return &Store{
		Name:        name,
		Assignments: map[uuid.UUID][]*assignment.Assignment{},
	}
}

// AssignItem assigns an item to the store with a given unit and quantity.
func (store *Store) AssignItem(
	item *item.Item,
	unit assignment.Unit,
	quantity assignment.Quantity,
) *assignment.Assignment {
	assignedItem := assignment.NewAssignment(item, unit, quantity)

	store.Assignments[item.Uid] = append(store.Assignments[item.Uid], assignedItem)

	return assignedItem
}

// GetAvailableQuantityForItem calculate the available quantity of an item
func (store *Store) GetAvailableQuantityForItem(item *item.Item) assignment.Quantity {
	var availableQuantity assignment.Quantity

	for _, itemAssignment := range store.Assignments[item.Uid] {
		availableQuantity += itemAssignment.GetAvailableQuantity()
	}

	return availableQuantity
}

// Uptake updates the store by taking a certain quantity of an item.
// If the available quantity is less than the requested quantity, it returns an error indicating the insufficient quantity.
func (store *Store) Uptake(item *item.Item, quantity assignment.Quantity) (assignment.Quantity, error) {
	availableQuantity := store.GetAvailableQuantityForItem(item)

	if availableQuantity < quantity {
		return 0, assignment.ErrorNotEnoughQuantity{Available: availableQuantity, Wanted: quantity}
	}

	assignments := store.Assignments[item.Uid]
	nbAssignments := len(assignments)
	remainingQuantityToUptake := quantity

	for index := 0; index < nbAssignments && remainingQuantityToUptake > 0; index++ {
		availableQuantityForAssignment := assignments[index].GetAvailableQuantity()
		quantityToUptake := remainingQuantityToUptake

		if availableQuantityForAssignment < quantityToUptake {
			quantityToUptake = availableQuantityForAssignment
		}

		_, err := assignments[index].Uptake(quantityToUptake)

		if err != nil {
			return store.GetAvailableQuantityForItem(item), err
		}

		remainingQuantityToUptake -= quantityToUptake
	}

	return store.GetAvailableQuantityForItem(item), nil
}
