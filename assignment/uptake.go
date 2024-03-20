package assignment

type Uptake struct {
	Quantity Quantity
}

func NewUptake(quantity Quantity) Uptake {
	return Uptake{
		Quantity: quantity,
	}
}
