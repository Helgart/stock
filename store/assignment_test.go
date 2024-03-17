package store

import (
	"github.com/google/uuid"
	"testing"
)

func TestAssignmentCreation(t *testing.T) {
	tests := []struct {
		item     Item
		unit     Unit
		quantity Quantity
	}{
		{NewItem("pasta"), "box", 1},
		{NewItem("rice"), "gram", 500},
	}

	for _, fixture := range tests {
		item := NewAssignment(fixture.item, fixture.unit, fixture.quantity)

		if item.Item != fixture.item {
			t.Errorf("Expected name %s, but got %s", fixture.item.Name, item.Item.Name)
		}
		if item.Uid == uuid.Nil {
			t.Errorf("Expected non-nil UUID, but got nil")
		}
		if item.Unit != fixture.unit {
			t.Errorf("Expected unit %s, but got %s", fixture.unit, item.Unit)
		}
		if item.Quantity != fixture.quantity {
			t.Errorf("Expected quantity %d, but got %d", fixture.quantity, item.Quantity)
		}
	}
}
