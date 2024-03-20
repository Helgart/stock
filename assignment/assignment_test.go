package assignment

import (
	"stock.ngmengineering.fr/stock/item"
	"testing"
)

func TestAssignmentCreation(t *testing.T) {
	tests := []struct {
		itemName string
		unit     Unit
		quantity Quantity
	}{
		{"pasta", "box", 1},
		{"rice", "gram", 500},
	}

	for _, fixture := range tests {
		itemToAssign, err := item.NewItem(fixture.itemName)

		if err != nil {
			t.Errorf("Error creating item: %v", err)
			continue
		}

		itemAssignment := NewAssignment(itemToAssign, fixture.unit, fixture.quantity)

		if itemAssignment.Item != itemToAssign {
			t.Errorf("Expected name %v, but got %v", itemToAssign, itemAssignment.Item)
		}
		if itemAssignment.Unit != fixture.unit {
			t.Errorf("Expected unit %s, but got %s", fixture.unit, itemAssignment.Unit)
		}
		if itemAssignment.Quantity != fixture.quantity {
			t.Errorf("Expected quantity %d, but got %d", fixture.quantity, itemAssignment.Quantity)
		}
	}
}
