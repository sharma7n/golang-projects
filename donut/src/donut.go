package src

// Shape ...
type Shape int

const (
	// Ring ...
	Ring Shape = 0

	// Hole ...
	Hole Shape = 1
)

// Donut ...
type Donut struct {
	Shape Shape
}