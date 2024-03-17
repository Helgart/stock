package assignment

type Uptake struct {
	Assignment Assignment
	Quantity   Quantity
}

func NewUptake(assignment Assignment, quantity Quantity) Uptake {
	return Uptake{
		Assignment: assignment,
		Quantity:   quantity,
	}
}
