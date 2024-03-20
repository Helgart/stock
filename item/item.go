package item

import (
	"github.com/google/uuid"
)

type Item struct {
	Uid  uuid.UUID
	Name string
}

func NewItem(name string) (Item, error) {
	uid, err := uuid.NewV7()

	if err != nil {
		return Item{}, err
	}

	return Item{
			Uid:  uid,
			Name: name,
		},
		err
}
