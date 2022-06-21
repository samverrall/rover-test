package direction

import (
	"fmt"
)

type Direction struct {
	Value string
	Right func() *Direction
	Left  func() *Direction
}

// New initialises a new Direction and returns a pointer
func New(value string) (*Direction, error) {
	out := &Direction{}

	switch value {
	case VNorth.String():
		return out.North()(), nil
	case VWest.String():
		return out.West()(), nil
	case VEast.String():
		return out.East()(), nil
	case VSouth.String():
		return out.South()(), nil
	default:
		return nil, fmt.Errorf("invalid direction supplied to New: %s", value)
	}
}

// These methods are indirectly tested within the `parse` package.
func (d *Direction) North() func() *Direction {
	return func() *Direction {
		return &Direction{
			Value: VNorth.String(),
			Right: d.East(),
			Left:  d.West(),
		}
	}
}

func (d *Direction) West() func() *Direction {
	return func() *Direction {
		return &Direction{
			Value: VWest.String(),
			Right: d.North(),
			Left:  d.South(),
		}
	}
}

func (d *Direction) South() func() *Direction {
	return func() *Direction {
		return &Direction{
			Value: VSouth.String(),
			Right: d.West(),
			Left:  d.East(),
		}
	}
}

func (d *Direction) East() func() *Direction {
	return func() *Direction {
		return &Direction{
			Value: VEast.String(),
			Right: d.South(),
			Left:  d.North(),
		}
	}
}

// IsDirection validates if a direction is valid,
// returning true if valid.
func (d *Direction) IsDirection(vd VehicleDirection) bool {
	return d.Value == vd.String()
}
