package rover

import (
	"testing"

	"github.com/samverrall/rover-test/coordinates"
	"github.com/samverrall/rover-test/direction"
	"github.com/samverrall/rover-test/plateau"
	"github.com/stretchr/testify/require"
)

func TestRover_Navigate(t *testing.T) {
	const (
		invalidCommand = "ABC"
	)

	type result struct {
		direction string
		xStart    int
		yStart    int
	}

	type args struct {
		direction string
		commands  string
		xStart    int
		yStart    int
		grid      *plateau.Plateau
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    result
	}{
		{
			name: "Empty commands supplied returns an error",
			args: args{
				direction: direction.VNorth.String(),
				commands:  "",
				xStart:    2,
				yStart:    2,
				grid:      plateau.New(5, 5),
			},
			wantErr: true,
		},
		{
			name: "Empty commands supplied returns an error",
			args: args{
				direction: direction.VNorth.String(),
				commands:  invalidCommand,
				xStart:    2,
				yStart:    2,
				grid:      plateau.New(5, 5),
			},
			wantErr: true,
		},
		{
			name: "Rotate rover right, expect East",
			args: args{
				direction: direction.VNorth.String(),
				commands:  "R",
				xStart:    2,
				yStart:    2,
				grid:      plateau.New(5, 5),
			},
			wantErr: false,
			want: result{
				direction: direction.VEast.String(),
				xStart:    2,
				yStart:    2,
			},
		},
		{
			name: "Rotate rover left, expect West",
			args: args{
				direction: direction.VNorth.String(),
				commands:  "L",
				xStart:    2,
				yStart:    2,
				grid:      plateau.New(5, 5),
			},
			wantErr: false,
			want: result{
				direction: direction.VWest.String(),
				xStart:    2,
				yStart:    2,
			},
		},
		{
			name: "Rotate rover left twice, expect south",
			args: args{
				direction: direction.VNorth.String(),
				commands:  "LL",
				xStart:    2,
				yStart:    2,
				grid:      plateau.New(5, 5),
			},
			wantErr: false,
			want: result{
				direction: direction.VSouth.String(),
				xStart:    2,
				yStart:    2,
			},
		},
		{
			name: "Rotate rover left twice, expect south",
			args: args{
				direction: direction.VNorth.String(),
				commands:  "LL",
				xStart:    2,
				yStart:    2,
				grid:      plateau.New(5, 5),
			},
			wantErr: false,
			want: result{
				direction: direction.VSouth.String(),
				xStart:    2,
				yStart:    2,
			},
		},
		{
			name: "Rotate rover left and move, expect west and x decreased",
			args: args{
				direction: direction.VNorth.String(),
				commands:  "LM",
				xStart:    2,
				yStart:    2,
				grid:      plateau.New(5, 5),
			},
			wantErr: false,
			want: result{
				direction: direction.VWest.String(),
				xStart:    1,
				yStart:    2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			direction, err := direction.New(tt.args.direction)
			require.NoError(t, err, "failed to create direction: %v", err)

			r := &Rover{
				Direction:  direction,
				Coordinate: coordinates.New(tt.args.xStart, tt.args.yStart),
			}
			err = r.Navigate(tt.args.commands, tt.args.grid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rover.Navigate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				require.Equal(t, tt.want.direction, r.Direction.Value, "want direction: %s, got: %s", tt.want.direction, r.Direction.Value)
				require.Equal(t, tt.want.xStart, r.Coordinate.GetX(), "want x: %s, got: %s", tt.want.xStart, r.Coordinate.GetX())
				require.Equal(t, tt.want.yStart, r.Coordinate.GetY(), "want y: %s, got: %s", tt.want.yStart, r.Coordinate.GetY())
			}
		})
	}
}

func TestRover_Position(t *testing.T) {
	type args struct {
		xStart    int
		yStart    int
		direction string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Returns rover position string",
			args: args{
				xStart:    4,
				yStart:    2,
				direction: "S",
			},
			want: "4 2 S",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			direction, err := direction.New(tt.args.direction)
			require.NoError(t, err, "failed to create direction: %v", err)

			r := &Rover{
				Direction:  direction,
				Coordinate: coordinates.New(tt.args.xStart, tt.args.yStart),
			}
			if got := r.Position(); got != tt.want {
				t.Errorf("Rover.Position() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	const (
		invalidDirectionString = "B"
	)

	type args struct {
		in RoverInArgs
	}
	tests := []struct {
		name    string
		args    args
		want    *Rover
		wantErr bool
	}{
		{
			name: "New successfully returns rover",
			args: args{
				in: RoverInArgs{
					Position: coordinates.Coordinate{
						X: 5,
						Y: 5,
					},
					DirectionString: "N",
				},
			},
			want: &Rover{
				Direction:  &direction.Direction{Value: direction.VNorth.String()},
				Coordinate: coordinates.New(5, 5),
			},
			wantErr: false,
		},
		{
			name: "New returns an error, invalid direction string",
			args: args{
				in: RoverInArgs{
					Position: coordinates.Coordinate{
						X: 5,
						Y: 5,
					},
					DirectionString: invalidDirectionString,
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil {
				require.EqualValues(t, got.Coordinate.GetX(), tt.want.Coordinate.GetX())
				require.EqualValues(t, got.Coordinate.GetY(), tt.want.Coordinate.GetY())
				require.EqualValues(t, got.Direction.Value, tt.want.Direction.Value)
			}
		})
	}
}
