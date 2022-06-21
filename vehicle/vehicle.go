package vehicle

import "github.com/samverrall/rover-test/plateau"

// Vehicle represents the required implementation for a vehicle.
type Vehicle interface {
	Position() string
	Navigate(commands string, grid *plateau.Plateau) error
}
