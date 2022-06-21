package coordinates

import (
	"reflect"
	"testing"

	"github.com/samverrall/rover-test/direction"
	"github.com/samverrall/rover-test/plateau"
	"github.com/stretchr/testify/require"
)

func TestCoordinate_Next(t *testing.T) {
	directionNorth, err := direction.New(direction.VNorth.String())
	require.NoError(t, err)

	directionSouth, err := direction.New(direction.VSouth.String())
	require.NoError(t, err)

	directionWest, err := direction.New(direction.VWest.String())
	require.NoError(t, err)

	directionEast, err := direction.New(direction.VEast.String())
	require.NoError(t, err)

	type fields struct {
		X int
		Y int
	}
	type args struct {
		curDirection *direction.Direction
		grid         *plateau.Plateau
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Coordinate
	}{
		{
			name: "Successfully moves north",
			fields: fields{
				X: 2,
				Y: 1,
			},
			args: args{
				curDirection: directionNorth,
				grid:         plateau.New(5, 5),
			},
			want: &Coordinate{
				X: 2,
				Y: 2,
			},
		},
		{
			name: "Successfully moves south",
			fields: fields{
				X: 2,
				Y: 2,
			},
			args: args{
				curDirection: directionSouth,
				grid:         plateau.New(5, 5),
			},
			want: &Coordinate{
				X: 2,
				Y: 1,
			},
		},
		{
			name: "Successfully moves west",
			fields: fields{
				X: 2,
				Y: 2,
			},
			args: args{
				curDirection: directionWest,
				grid:         plateau.New(5, 5),
			},
			want: &Coordinate{
				X: 1,
				Y: 2,
			},
		},
		{
			name: "Successfully moves east",
			fields: fields{
				X: 2,
				Y: 2,
			},
			args: args{
				curDirection: directionEast,
				grid:         plateau.New(5, 5),
			},
			want: &Coordinate{
				X: 3,
				Y: 2,
			},
		},
		{
			name: "Returns the same, as Y is already at max",
			fields: fields{
				X: 5,
				Y: 5,
			},
			args: args{
				curDirection: directionNorth,
				grid:         plateau.New(5, 5),
			},
			want: &Coordinate{
				X: 5,
				Y: 5,
			},
		},
		{
			name: "Returns the same, as X is already at max",
			fields: fields{
				X: 5,
				Y: 5,
			},
			args: args{
				curDirection: directionEast,
				grid:         plateau.New(5, 5),
			},
			want: &Coordinate{
				X: 5,
				Y: 5,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Coordinate{
				X: tt.fields.X,
				Y: tt.fields.Y,
			}
			if got := c.Next(tt.args.curDirection, tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Coordinate.Next() = %v, want %v", got, tt.want)
			}
		})
	}
}
