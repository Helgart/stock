package item

import (
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
		item, err := NewItem(fixture)

		s.Require().NoError(err)
		s.Equal(fixture, item.Name)
		s.NotEqual(item.Uid, uuid.Nil)
	}
}
