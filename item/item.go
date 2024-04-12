package item

import (
	"github.com/google/uuid"
)

type UIDProvider func() (uuid.UUID, error)

// Item represents an item with a Uid and a Name.
type Item struct {
	Uid  uuid.UUID
	Name string
}

// NewItem creates a new item with a unique identifier and the given name.
// Returns the new item and an error, if there was an error generating the unique identifier.
func NewItem(name string, provider UIDProvider) (*Item, error) {
	uid, err := provider()

	if err != nil {
		return nil, err
	}

	return &Item{Uid: uid, Name: name}, err
}
