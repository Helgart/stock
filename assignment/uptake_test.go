package assignment

import (
	"testing"
)

func TestUptakeCreation(t *testing.T) {
	tests := []struct {
		assignment Assignment
		quantity   Quantity
	}{
		{NewValidAssignmentFixture(5), 5},
		{NewValidAssignmentFixture(5), 1},
	}

	for _, fixture := range tests {
		uptake := NewUptake(fixture.assignment, fixture.quantity)

		if uptake.Assignment != fixture.assignment {
			t.Errorf("Expected assignment %v, but got %v", fixture.assignment, uptake.Assignment)
		}
		if uptake.Quantity != fixture.quantity {
			t.Errorf("Expected quantity %d, but got %d", fixture.quantity, uptake.Quantity)
		}
	}
}
