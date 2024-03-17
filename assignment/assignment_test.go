package assignment

import (
	"stock.ngmengineering.fr/stock/item"
	"stock.ngmengineering.fr/stock/store"
	"testing"
)

func TestAssignmentCreation(t *testing.T) {
	tests := []struct {
		item     item.Item
		store    store.Store
		unit     Unit
		quantity Quantity
	}{
		{item.NewItem("pasta"), store.NewStore("cupboard"), "box", 1},
		{item.NewItem("rice"), store.NewStore("cupboard"), "gram", 500},
	}

	for _, fixture := range tests {
		itemAssignment := NewAssignment(fixture.item, fixture.store, fixture.unit, fixture.quantity)

		if itemAssignment.Item != fixture.item {
			t.Errorf("Expected name %s, but got %s", fixture.item.Name, itemAssignment.Item.Name)
		}
		if itemAssignment.Unit != fixture.unit {
			t.Errorf("Expected unit %s, but got %s", fixture.unit, itemAssignment.Unit)
		}
		if itemAssignment.Quantity != fixture.quantity {
			t.Errorf("Expected quantity %d, but got %d", fixture.quantity, itemAssignment.Quantity)
		}
	}
}

func NewValidAssignmentFixture(quantity Quantity) Assignment {
	return NewAssignment(item.NewItem("pasta"), store.NewStore("cupboard"), "box", quantity)
}
