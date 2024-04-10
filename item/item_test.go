package item_test

import (
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

	for _, fixture := range tests {
		fixture, err := item.NewItem(fixture)

		s.Require().NoError(err)
		s.Equal(fixture, fixture.Name)
		s.NotEqual(fixture.Uid, uuid.Nil)
	}
}
