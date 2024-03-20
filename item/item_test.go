package item

import (
	"fmt"
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
		item, err := NewItem(fixture)

		if item.Name != fixture {
			t.Errorf("Expected name %s, but got %s", fixture, item.Name)
		}
		if item.Uid == uuid.Nil {
			t.Errorf("Expected non-nil UUID, but got nil")
		}
		if err != nil {
			t.Errorf(fmt.Sprintf("Expected no error, but got one : %s", err.Error()))
		}
	}
}
