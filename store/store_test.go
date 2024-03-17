package store

import (
	"github.com/google/uuid"
	"testing"
)

func TestStoreCreation(t *testing.T) {
	tests := []string{
		"fridge",
		"box",
		"freezer",
	}

	for _, fixture := range tests {
		item := NewStore(fixture)

		if item.Name != fixture {
			t.Errorf("Expected name %s, but got %s", fixture, item.Name)
		}
		if item.Uid == uuid.Nil {
			t.Errorf("Expected non-nil UUID, but got nil")
		}
	}
}
