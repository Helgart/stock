package assignment

import (
	"fmt"
	"stock.ngmengineering.fr/stock/item"
)

// notEnoughQuantityErrorMessage is a constant string used as an error message when a quantity is not sufficient.
// It provides a formatted string with the available quantity and the desired quantity.
// Example usage:
// ```go
//
//	func (notEnoughQuantity NotEnoughQuantity) Error() string {
//		return fmt.Sprintf(notEnoughQuantityErrorMessage, notEnoughQuantity.Available, notEnoughQuantity.Wanted)
//	}
//
// ```
// For example, if `Available` is 5 and `Wanted` is 10, the error message will be "Not enough quantity. Available: 5, Wanted: 10".
const notEnoughQuantityErrorMessage = "Not enough quantity. Available: %d, Wanted: %d"

// Unit represents a unit of measurement for an Assignment.
type Unit string

// Quantity represents a numerical value used for measuring a certain quantity in the context of an Assignment.
type Quantity uint

// Assignment represents an assignment of an item, along with its unit, quantity, and uptakes.
type Assignment struct {
	Item     item.Item
	Unit     Unit
	Quantity Quantity
	Uptakes  []Uptake
}

// NotEnoughQuantity represents an error indicating that there is not enough quantity available for a specific operation.
type NotEnoughQuantity struct {
	Available Quantity
	Wanted    Quantity
}

// Error returns the error message for the NotEnoughQuantity error type.
// It formats the error message using the values of the Available and Wanted properties.
func (notEnoughQuantity NotEnoughQuantity) Error() string {
	return fmt.Sprintf(notEnoughQuantityErrorMessage, notEnoughQuantity.Available, notEnoughQuantity.Wanted)
}

// NewAssignment creates a new Assignment with the specified item, unit, and quantity.
func NewAssignment(
	item item.Item,
	unit Unit,
	quantity Quantity,
) Assignment {
	return Assignment{
		Item:     item,
		Unit:     unit,
		Quantity: quantity,
	}
}

// GetAvailableQuantity calculates the available quantity of an assignment by subtracting the uptake quantities from the assignment quantity.
// It returns the available quantity.
//
// Example:
//
//	assignment := Assignment{
//	    Quantity: 10,
//	    Uptakes:  []Uptake{
//	        {Quantity: 2},
//	        {Quantity: 3},
//	    },
//	}
//	availableQuantity := assignment.GetAvailableQuantity() // availableQuantity = 5
func (assignment *Assignment) GetAvailableQuantity() Quantity {
	availableQuantity := assignment.Quantity

	for _, uptake := range assignment.Uptakes {
		availableQuantity -= uptake.Quantity
	}

	return availableQuantity
}

// Uptake updates the Assignment by subtracting the given quantity from the available quantity.
// If the available quantity is less than the given quantity, it returns a NotEnoughQuantity error.
func (assignment *Assignment) Uptake(quantity Quantity) (Quantity, error) {
	availableQuantity := assignment.GetAvailableQuantity()

	if availableQuantity < quantity {
		return 0, NotEnoughQuantity{availableQuantity, quantity}
	}

	assignment.Uptakes = append(assignment.Uptakes, NewUptake(quantity))

	return assignment.GetAvailableQuantity(), nil
}
