package item_test

import (
	"errors"
	"github.com/Helgart/stock/item"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"testing"
)

type itemTestSuite struct {
	suite.Suite
}

func TestItemTestSuite(t *testing.T) {
	suite.Run(t, new(itemTestSuite))
}

func (s *itemTestSuite) TestItemCreation() {
	tests := []string{
		"pasta",
		"ham",
		"bread",
	}
	generatedUuid, err := uuid.NewUUID()

	s.Require().NoError(err)

	provider := func() (uuid.UUID, error) {
		return generatedUuid, nil
	}

	for _, fixture := range tests {
		fixtureItem, err := item.NewItem(fixture, provider)

		s.Require().NoError(err)
		s.Equal(fixture, fixtureItem.Name)
		s.Equal(fixtureItem.Uid, generatedUuid)
	}
}

func (s *itemTestSuite) TestItemCreationWithMockError() {
	provider := func() (uuid.UUID, error) {
		return uuid.UUID{}, errors.New("mock error")
	}

	_, err := item.NewItem("Bread", provider)

	s.Require().Error(err)
}
