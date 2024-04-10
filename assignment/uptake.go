package assignment

// Uptake represents the quantity of an item taken or consumed.
type Uptake struct {
	Quantity Quantity
}

// NewUptake creates a new Uptake with the specified quantity.
func NewUptake(quantity Quantity) *Uptake {
	return &Uptake{
		Quantity: quantity,
	}
}
