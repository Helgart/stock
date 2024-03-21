package assignment

import (
	"fmt"
	"stock.ngmengineering.fr/stock/item"
)

const notEnoughQuantityErrorMessage = "Not enough quantity. Available: %d, Wanted: %d"

type Unit string
type Quantity uint

type Assignment struct {
	Item     item.Item
	Unit     Unit
	Quantity Quantity
	Uptakes  []Uptake
}

type NotEnoughQuantity struct {
	Available Quantity
	Wanted    Quantity
}

func (notEnoughQuantity NotEnoughQuantity) Error() string {
	return fmt.Sprintf(notEnoughQuantityErrorMessage, notEnoughQuantity.Available, notEnoughQuantity.Wanted)
}

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

func (assignment *Assignment) GetAvailableQuantity() Quantity {
	availableQuantity := assignment.Quantity

	for _, uptake := range assignment.Uptakes {
		availableQuantity -= uptake.Quantity
	}

	return availableQuantity
}

func (assignment *Assignment) Uptake(quantity Quantity) (Quantity, error) {
	availableQuantity := assignment.GetAvailableQuantity()

	if availableQuantity < quantity {
		return 0, NotEnoughQuantity{availableQuantity, quantity}
	}

	assignment.Uptakes = append(assignment.Uptakes, NewUptake(quantity))

	return assignment.GetAvailableQuantity(), nil
}
