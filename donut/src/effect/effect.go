package effect

import (
	"app/gen/donut"
)

// GetDonutList produces a DonutList containing all donuts currently in store.
type GetDonutList interface {
	GetDonutList() (donut.DonutList, error)
}

// AddDonut adds a Donut to the store.
type AddDonut interface {
	AddDonut(donut donut.Donut) error
}