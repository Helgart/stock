package assignment_test

import (
	"github.com/Helgart/stock/assignment"
	"github.com/Helgart/stock/item"
	"github.com/stretchr/testify/suite"
	"testing"
)

type assignmentTestSuite struct {
	suite.Suite
}

func TestAssignmentTestSuite(t *testing.T) {
	suite.Run(t, new(assignmentTestSuite))
}

func (s *assignmentTestSuite) TestAssignment_Creation() {
	tests := []struct {
		name     string
		itemName string
		unit     assignment.Unit
		quantity assignment.Quantity
	}{
		{"smallQuantity", "pasta", "box", 1},
		{"biggerQuantity", "rice", "gram", 500},
	}

	for _, fixture := range tests {
		s.Run(fixture.name, func() {
			itemToAssign, err := item.NewItem(fixture.itemName)
			itemAssignment := assignment.NewAssignment(itemToAssign, fixture.unit, fixture.quantity)

			s.Require().NoError(err)
			s.Equal(itemToAssign, itemAssignment.Item)
			s.Equal(fixture.unit, itemAssignment.Unit)
			s.Equal(fixture.quantity, itemAssignment.Quantity)
		})
	}
}

func (s *assignmentTestSuite) TestAssignment_GetAvailableQuantity() {
	tests := []struct {
		name     string
		itemName string
		unit     assignment.Unit
		quantity assignment.Quantity
		uptake   assignment.Quantity
	}{
		{"uptakeAll", "pasta", "box", 1, 1},
		{"uptakeHalf", "rice", "gram", 500, 250},
	}

	for _, fixture := range tests {
		s.Run(fixture.name, func() {
			itemToAssign, err := item.NewItem(fixture.itemName)

			s.Require().NoError(err)

			itemAssignment := assignment.NewAssignment(itemToAssign, fixture.unit, fixture.quantity)
			itemUptake := assignment.NewUptake(fixture.uptake)

			itemAssignment.Uptakes = append(itemAssignment.Uptakes, itemUptake)

			currentQuantity := itemAssignment.GetAvailableQuantity()
			expectedQuantity := fixture.quantity - fixture.uptake

			s.Equal(expectedQuantity, currentQuantity)
		})
	}
}

func (s *assignmentTestSuite) TestAssignment_Uptake() {
	tests := []struct {
		name     string
		itemName string
		unit     assignment.Unit
		quantity assignment.Quantity
		uptake   assignment.Quantity
	}{
		{"uptakeAll", "pasta", "box", 1, 1},
		{"uptakeHalf", "rice", "gram", 500, 250},
		{"uptakeMoreThanAvailable", "floor", "gram", 1000, 1200},
	}

	for _, fixture := range tests {
		s.Run(fixture.name, func() {
			itemToAssign, err := item.NewItem(fixture.itemName)

			s.Require().NoError(err)

			itemAssignment := assignment.NewAssignment(itemToAssign, fixture.unit, fixture.quantity)

			currentQuantity, err := itemAssignment.Uptake(fixture.uptake)

			if fixture.uptake > fixture.quantity {
				s.Require().Error(err)
				s.IsType(&assignment.ErrorNotEnoughQuantity{}, err)
			} else {
				s.Require().NoError(err)
				s.Equal(fixture.quantity-fixture.uptake, currentQuantity)
			}
		})
	}
}
