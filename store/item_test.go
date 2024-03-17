package store

import (
	"github.com/google/uuid"
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
		if item.Uid == uuid.Nil {
			t.Errorf("Expected non-nil UUID, but got nil")
		}
	}
}
