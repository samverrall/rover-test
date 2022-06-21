package parser

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_splitLines(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  [][]string
	}{
		{
			name:  "Empty input supplied",
			input: "",
			want:  [][]string{},
		},
		{
			name: "Successfully splits lines",
			input: `
5 5

1 2 N

LMLMLMLMM

3 3 E

MMRMMRMRRM
`,
			want: [][]string{
				{"5 5"},
				{"1 2 N"},
				{"LMLMLMLMM"},
				{"3 3 E"},
				{"MMRMMRMRRM"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitLines(tt.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		in string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name:    "Empty string parsed",
			args:    args{},
			wantErr: true,
		},
		{
			name: "Invalid grid input supplied",
			args: args{
				in: `
				L 2
			
				1 2 N
			
				LMLMLMLMM
			
				3 3 E
			
				MMRMMRMRRM
				`,
			},
			wantErr: true,
		},
		{
			name: "No vehicles supplied returns an err",
			args: args{
				in: `
				5 5
			
				`,
			},
			wantErr: true,
		},
		{
			name: "Zero size grid supplied",
			args: args{
				in: `
				0 0
			
				1 2 N
			
				LMLMLMLMM
			
				3 3 E
			
				MMRMMRMRRM
				`,
			},
			wantErr: true,
		},
		{
			name: "Successfully parse",
			args: args{
				in: `
				5 5
			
				1 2 N
			
				LMLMLMLMM
			
				3 3 E
			
				MMRMMRMRRM
				`,
			},
			want:    []string{"1 3 N", "5 1 E"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil {
				for idx, vehicle := range got.Vicheles {
					require.Equal(t, vehicle.Position(), tt.want[idx])
				}
			}
		})
	}
}
