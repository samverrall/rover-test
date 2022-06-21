package rover

import (
	"errors"
	"fmt"
	"strings"

	"github.com/samverrall/rover-test/coordinates"
	"github.com/samverrall/rover-test/direction"
	"github.com/samverrall/rover-test/plateau"
)

const (
	Right = 'R'
	Left  = 'L'
	Move  = 'M'
)

type Rover struct {
	Direction  *direction.Direction
	Coordinate coordinates.Coorninator
}

type RoverInArgs struct {
	Position        coordinates.Coordinate
	DirectionString string
}

func New(in RoverInArgs) (*Rover, error) {
	direction, err := direction.New(in.DirectionString)
	if err != nil {
		return nil, err
	}

	return &Rover{
		// I purposely initialise direction struct here myself rather than use a `New()` initiator as there is no initiation values needed.
		Direction:  direction,
		Coordinate: coordinates.New(in.Position.X, in.Position.Y),
	}, nil
}

func (r *Rover) Position() string {
	return fmt.Sprintf("%d %d %s", r.Coordinate.GetX(), r.Coordinate.GetY(), r.Direction.Value)
}

// Navigate executes the instructions to move a vehicle.
// Invalid instructions will cause an error to be returned.
func (r *Rover) Navigate(commands string, grid *plateau.Plateau) error {
	if strings.TrimSpace(commands) == "" {
		return errors.New("no commands supplied, failed to naviate")
	}

	// Iterate through the commands
	for _, command := range commands {
		switch command {
		case Right:
			r.Direction = r.Direction.Right()
		case Left:
			r.Direction = r.Direction.Left()
		case Move:
			r.Coordinate = r.Coordinate.Next(r.Direction, grid)
		default:
			return fmt.Errorf("invalid command supplied, %v", command)
		}
	}

	return nil
}
