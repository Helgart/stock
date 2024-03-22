package assignment

import (
	"github.com/Helgart/stock/item"
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

func TestAssignment_GetAvailableQuantity(t *testing.T) {
	tests := []struct {
		itemName string
		unit     Unit
		quantity Quantity
		uptake   Quantity
	}{
		{"pasta", "box", 1, 1},
		{"rice", "gram", 500, 250},
	}

	for _, fixture := range tests {
		itemToAssign, err := item.NewItem(fixture.itemName)

		if err != nil {
			t.Errorf("Error creating item: %v", err)
			continue
		}

		itemAssignment := NewAssignment(itemToAssign, fixture.unit, fixture.quantity)
		itemUptake := NewUptake(fixture.uptake)

		itemAssignment.Uptakes = append(itemAssignment.Uptakes, itemUptake)

		currentQuantity := itemAssignment.GetAvailableQuantity()
		expectedQuantity := fixture.quantity - fixture.uptake

		if currentQuantity != fixture.quantity-fixture.uptake {
			t.Errorf("Expected %d quantity after uptake, obtained %d", expectedQuantity, currentQuantity)
		}
	}
}

func TestAssignment_Uptake(t *testing.T) {
	tests := []struct {
		itemName string
		unit     Unit
		quantity Quantity
		uptake   Quantity
	}{
		{"pasta", "box", 1, 1},
		{"rice", "gram", 500, 250},
		{"floor", "gram", 1000, 1200},
	}

	for _, fixture := range tests {
		itemToAssign, err := item.NewItem(fixture.itemName)

		if err != nil {
			t.Errorf("Error creating item: %v", err)
			continue
		}

		itemAssignment := NewAssignment(itemToAssign, fixture.unit, fixture.quantity)

		currentQuantity, err := itemAssignment.Uptake(fixture.uptake)

		if fixture.uptake > fixture.quantity && err == nil {
			t.Errorf("Expected error after uptake, none obtained")
		}
		if fixture.uptake > fixture.quantity && err != nil {
			if _, isValid := err.(NotEnoughQuantity); !isValid {
				t.Errorf("Expected NotEnoughQuantity error after uptake, %T obtained", err)
			}
		}
		if fixture.uptake <= fixture.quantity && currentQuantity != fixture.quantity-fixture.uptake {
			t.Errorf(
				"Expected %d quantity after uptake, obtained %d",
				fixture.quantity-fixture.uptake,
				currentQuantity,
			)
		}
	}
}
