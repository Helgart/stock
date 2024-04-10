package store_test

import (
	"github.com/Helgart/stock/assignment"
	"github.com/Helgart/stock/item"
	"github.com/Helgart/stock/store"
	"github.com/stretchr/testify/suite"
	"testing"
)

type storeTestSuite struct {
	suite.Suite
}

func TestStoreTestSuite(t *testing.T) {
	suite.Run(t, new(storeTestSuite))
}

func (s *storeTestSuite) TestStore_Creation() {
	tests := []string{
		"fridge",
		"box",
		"freezer",
	}

	for _, fixture := range tests {
		s.Run(fixture, func() {
			storeFixture := store.NewStore(fixture)

			s.Equal(fixture, storeFixture.Name)
			s.Empty(storeFixture.Assignments)
		})
	}
}

func (s *storeTestSuite) TestStore_AssignItem() {
	storeFixture := store.NewStore("fridge")
	tests := []struct {
		itemName string
		unit     assignment.Unit
		quantity assignment.Quantity
	}{
		{"pasta", "box", 5},
		{"drink", "bottle", 10},
	}
	expectedLength := 1

	for _, test := range tests {
		s.Run(test.itemName, func() {
			itemToAssign, err := item.NewItem(test.itemName)

			s.Require().NoError(err)

			currentAssignment := storeFixture.AssignItem(itemToAssign, test.unit, test.quantity)

			s.Len(storeFixture.Assignments, expectedLength)
			s.Equal(currentAssignment.Item, itemToAssign)
			s.Equal(currentAssignment.Quantity, test.quantity)
			s.Equal(currentAssignment.Unit, test.unit)

			expectedLength++
		})
	}
}

func (s *storeTestSuite) TestStore_GetAvailableQuantityForItem() {
	storeFixture := store.NewStore("fridge")
	tests := []struct {
		itemName string
		unit     assignment.Unit
		quantity assignment.Quantity
	}{
		{"pasta", "box", 5},
		{"drink", "bottle", 10},
	}

	for _, test := range tests {
		s.Run(test.itemName, func() {
			itemToAssign, err := item.NewItem(test.itemName)

			s.Require().NoError(err)

			storeFixture.AssignItem(itemToAssign, test.unit, test.quantity)
			availableQuantity := storeFixture.GetAvailableQuantityForItem(itemToAssign)

			s.Equal(test.quantity, availableQuantity)
		})
	}
}

func (s *storeTestSuite) TestStore_GetAvailableQuantityForItem_InvalidItem() {
	storeFixture := store.NewStore("fridge")
	theUnknownItem, err := item.NewItem("The Unknown")

	s.Require().NoError(err)

	availableQuantity := storeFixture.GetAvailableQuantityForItem(theUnknownItem)

	s.Empty(availableQuantity)
}

func (s *storeTestSuite) TestStore_Uptake_ValidQuantities() {
	storeFixture := store.NewStore("fridge")
	tests := []struct {
		itemName      string
		quantity      assignment.Quantity
		nbAssignments uint
		nbToUptake    assignment.Quantity
	}{
		{"uptakesAll", 5, 2, 10},
		{"updatesHalf", 4, 3, 6},
	}

	for _, test := range tests {
		s.Run(test.itemName, func() {
			itemToAssign, err := item.NewItem(test.itemName)

			s.Require().NoError(err)

			expectedRemainingQuantity := assignment.Quantity((uint(test.quantity) * test.nbAssignments) - uint(test.nbToUptake))

			for i := uint(0); i < test.nbAssignments; i++ {
				storeFixture.AssignItem(itemToAssign, "unit", test.quantity)
			}

			remainingQuantity, err := storeFixture.Uptake(itemToAssign, test.nbToUptake)

			s.Equal(expectedRemainingQuantity, remainingQuantity)
		})
	}
}

func (s *storeTestSuite) TestStore_Uptake_InvalidQuantities() {
	storeFixture := store.NewStore("fridge")
	tests := []struct {
		itemName      string
		quantity      assignment.Quantity
		nbAssignments uint
		nbToUptake    assignment.Quantity
	}{
		{"uptakesMoreThanAvailable", 5, 2, 15},
	}

	for _, test := range tests {
		s.Run(test.itemName, func() {
			itemToAssign, err := item.NewItem(test.itemName)

			for i := uint(0); i < test.nbAssignments; i++ {
				storeFixture.AssignItem(itemToAssign, "unit", test.quantity)
			}

			_, err = storeFixture.Uptake(itemToAssign, test.nbToUptake)

			s.Require().Error(err)
		})
	}
}
