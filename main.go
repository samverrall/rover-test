package main

import (
	"fmt"

	"github.com/samverrall/rover-test/parser"
)

func main() {
	input := `
	5 5

	1 2 N

	LMLMLMLMM

	3 3 E

	MMRMMRMRRM
	`

	data, err := parser.Parse(input)
	if err != nil {
		fmt.Printf("Failed to parse input: %v\n", err)
		return
	}

	for idx, vehicle := range data.Vicheles {
		fmt.Printf("Vehicle %d position output: %v\n", idx, vehicle.Position())
	}
}
