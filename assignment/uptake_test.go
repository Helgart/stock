package assignment

import "testing"

func TestUptakeCreation(t *testing.T) {
	tests := []Quantity{5, 1}

	for _, quantity := range tests {
		uptake := NewUptake(quantity)

		if uptake.Quantity != quantity {
			t.Errorf("Expected quantity %d, but got %d", quantity, uptake.Quantity)
		}
	}
}
