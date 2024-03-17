package store

import (
	"fmt"
	"github.com/google/uuid"
)

type Store struct {
	Uid  uuid.UUID
	Name string
}

func NewStore(name string) Store {
	uid, err := uuid.NewV7()

	if err != nil {
		panic(fmt.Sprintf(uuidGenerationError, err.Error()))
	}

	return Store{
		Uid:  uid,
		Name: name,
	}
}
