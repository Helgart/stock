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

func TestStoreAssignItem(t *testing.T) {
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

		currentAssignment := store.assignItem(itemToAssign, fixture.unit, fixture.quantity)
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
