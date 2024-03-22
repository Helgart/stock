package store

import (
	"stock.ngmengineering.fr/stock/assignment"
	"stock.ngmengineering.fr/stock/item"
	"testing"
)

func TestStoreCreation(t *testing.T) {
	tests := []string{
		"fridge",
		"box",
		"freezer",
	}

	for _, fixture := range tests {
		store := NewStore(fixture)

		if store.Name != fixture {
			t.Errorf("Expected name %s, but got %s", fixture, store.Name)
		}
		if len(store.Assignments) != 0 {
			t.Errorf("Expected empty assignments, but got %d assignments", len(store.Assignments))
		}
	}
}

func TestStore_AssignItem(t *testing.T) {
	store := NewStore("fridge")
	tests := []struct {
		itemName string
		unit     assignment.Unit
		quantity assignment.Quantity
	}{
		{"pasta", "box", 5},
		{"drink", "bottle", 10},
	}
	expectedLength := 1

	for _, fixture := range tests {
		itemToAssign, err := item.NewItem(fixture.itemName)

		if err != nil {
			t.Errorf("Expected no error on item creation, but got %s error", err.Error())
		}

		currentAssignment := store.AssignItem(itemToAssign, fixture.unit, fixture.quantity)
		nbAssignedItems := len(store.Assignments)

		if nbAssignedItems != expectedLength {
			t.Errorf("Expected %d assigned items, but got %d assignments", expectedLength, nbAssignedItems)
		}
		if currentAssignment.Item != itemToAssign {
			t.Errorf("Expected item %v, but got %v", itemToAssign, currentAssignment)
		}
		if currentAssignment.Quantity != fixture.quantity {
			t.Errorf("Expected quantity %d, but got %d", fixture.quantity, currentAssignment.Quantity)
		}
		if currentAssignment.Unit != fixture.unit {
			t.Errorf("Expected unit %s, but got %s", fixture.unit, currentAssignment.Unit)
		}

		expectedLength++
	}
}

func TestStore_GetAvailableQuantityForItem(t *testing.T) {
	store := NewStore("fridge")
	tests := []struct {
		itemName string
		unit     assignment.Unit
		quantity assignment.Quantity
	}{
		{"pasta", "box", 5},
		{"drink", "bottle", 10},
	}

	for _, fixture := range tests {
		itemToAssign, err := item.NewItem(fixture.itemName)

		if err != nil {
			t.Errorf("Expected no error on item creation, but got %s error", err.Error())
		}

		store.AssignItem(itemToAssign, fixture.unit, fixture.quantity)
		availableQuantity := store.GetAvailableQuantityForItem(itemToAssign)

		if availableQuantity != fixture.quantity {
			t.Errorf("Expected available quantity %d, but got %d", fixture.quantity, availableQuantity)
		}
	}
}

func TestStore_GetAvailableQuantityForItem_InvalidItem(t *testing.T) {
	store := NewStore("fridge")
	theUnknownItem, err := item.NewItem("The Unknown")

	if err != nil {
		t.Errorf("Expected no error on item creation, but got %s error", err.Error())
	}

	availableQuantity := store.GetAvailableQuantityForItem(theUnknownItem)

	if availableQuantity != 0 {
		t.Errorf("Expected available quantity 0, but got %d", availableQuantity)
	}
}

func TestStore_Uptake_ValidQuantities(t *testing.T) {
	store := NewStore("fridge")
	tests := []struct {
		quantity      assignment.Quantity
		nbAssignments uint
		nbToUptake    assignment.Quantity
	}{
		{5, 2, 10},
		{4, 3, 6},
	}

	for _, fixture := range tests {
		itemToAssign, err := item.NewItem("myItem")
		expectedRemainingQuantity := assignment.Quantity(
			(uint(fixture.quantity) * fixture.nbAssignments) - uint(fixture.nbToUptake),
		)

		for i := uint(0); i < fixture.nbAssignments; i++ {
			store.AssignItem(itemToAssign, "unit", fixture.quantity)
		}

		remainingQuantity, err := store.Uptake(itemToAssign, fixture.nbToUptake)

		if err != nil {
			t.Errorf("Expected no error on uptake, but got %s error", err.Error())
		}

		if remainingQuantity != expectedRemainingQuantity {
			t.Errorf("Expected remaining quantity %d, but got %d", expectedRemainingQuantity, remainingQuantity)
		}
	}
}

func TestStore_Uptake_InvalidQuantities(t *testing.T) {
	store := NewStore("fridge")
	tests := []struct {
		quantity      assignment.Quantity
		nbAssignments uint
		nbToUptake    assignment.Quantity
	}{
		{5, 2, 15},
	}

	for _, fixture := range tests {
		itemToAssign, err := item.NewItem("myItem")

		for i := uint(0); i < fixture.nbAssignments; i++ {
			store.AssignItem(itemToAssign, "unit", fixture.quantity)
		}

		_, err = store.Uptake(itemToAssign, fixture.nbToUptake)

		if err == nil {
			t.Errorf("Expected error on uptake, but got no error")
		}
	}
}
