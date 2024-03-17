package store

import (
	"fmt"
	"github.com/google/uuid"
)

type Unit string
type Quantity int

type Assignment struct {
	Uid      uuid.UUID
	Item     Item
	Unit     Unit
	Quantity Quantity
}

func NewAssignment(item Item, unit Unit, quantity Quantity) Assignment {
	uid, err := uuid.NewV7()

	if err != nil {
		panic(fmt.Sprintf(uuidGenerationError, err.Error()))
	}

	return Assignment{
		Uid:      uid,
		Item:     item,
		Unit:     unit,
		Quantity: quantity,
	}
}
