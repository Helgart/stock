package item

import (
	"testing"
)

func TestItemCreation(t *testing.T) {
	tests := []string{
		"pasta",
		"ham",
		"bread",
	}

	for _, fixture := range tests {
		item := NewItem(fixture)

		if item.Name != fixture {
			t.Errorf("Expected name %s, but got %s", fixture, item.Name)
		}
	}
}
