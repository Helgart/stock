package assignment_test

import (
	"github.com/Helgart/stock/assignment"
	"github.com/stretchr/testify/suite"
	"testing"
)

type uptakeTestSuite struct {
	suite.Suite
}

func TestUptakeTestSuite(t *testing.T) {
	suite.Run(t, new(uptakeTestSuite))
}

func (s *uptakeTestSuite) TestUptakeCreation() {
	tests := []struct {
		name     string
		quantity assignment.Quantity
	}{
		{"quantity", 5},
		{"smallQuantity", 1},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			uptake := assignment.NewUptake(test.quantity)
			s.Equal(test.quantity, uptake.Quantity)
		})
	}
}
