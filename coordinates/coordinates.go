package coordinates

import (
	"github.com/samverrall/rover-test/direction"
	"github.com/samverrall/rover-test/plateau"
)

type Coorninator interface {
	Next(curDirection *direction.Direction, grid *plateau.Plateau) *Coordinate

	GetX() int
	GetY() int
}

type Coordinate struct {
	X int
	Y int
}

func New(x int, y int) *Coordinate {
	return &Coordinate{
		X: x,
		Y: y,
	}
}

// Next moves a vehicle to by the supplied direction and returns the new vehicle coordinates.
func (c *Coordinate) Next(curDirection *direction.Direction, grid *plateau.Plateau) *Coordinate {
	switch {
	case curDirection.IsDirection(direction.VNorth) && c.Y < grid.GetHeight():
		return &Coordinate{
			X: c.X,
			Y: c.Y + 1,
		}
	case curDirection.IsDirection(direction.VSouth) && c.Y > 0:
		return &Coordinate{
			X: c.X,
			Y: c.Y - 1,
		}
	case curDirection.IsDirection(direction.VWest) && c.X > 0:
		return &Coordinate{
			X: c.X - 1,
			Y: c.Y,
		}
	case curDirection.IsDirection(direction.VEast) && c.X < grid.GetWidth():
		return &Coordinate{
			X: c.X + 1,
			Y: c.Y,
		}
	default:
		return c
	}
}

func (c *Coordinate) GetX() int {
	return c.X
}

func (c *Coordinate) GetY() int {
	return c.Y
}
