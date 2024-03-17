package store

import (
	"fmt"
	"github.com/google/uuid"
)

const uuidGenerationError = "Couldn't generate an UUID, current error is : %s"

type Item struct {
	Uid  uuid.UUID
	Name string
}

func NewItem(name string) Item {
	uid, err := uuid.NewV7()

	if err != nil {
		panic(fmt.Sprintf(uuidGenerationError, err.Error()))
	}

	return Item{
		Uid:  uid,
		Name: name,
	}
}
