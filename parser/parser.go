package parser

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/samverrall/rover-test/coordinates"
	"github.com/samverrall/rover-test/plateau"
	"github.com/samverrall/rover-test/vehicle"
	"github.com/samverrall/rover-test/vehicle/rover"
)

const (
	expectedStartingPositions  = 3
	expectedVehicleStringParts = 2

	gridFieldsCount = 2
)

type startingPositions struct {
	X         int
	Y         int
	Direction string
}

type ParserOut struct {
	Vicheles []vehicle.Vehicle
	Grid     *plateau.Plateau
}

func Parse(in string) (*ParserOut, error) {
	if strings.TrimSpace(in) == "" {
		return nil, errors.New("no input supplied")
	}

	parserOut := ParserOut{}
	inputSlices := splitLines(in)

	// Parse the first line which we know is the plateau size.
	width, height, err := validateGridSize(inputSlices[0])
	if err != nil {
		return nil, err
	}
	parserOut.Grid = plateau.New(width, height)

	// Parse all the lines after the plateau size.
	getVehiclesFromLines, err := getVehiclesFromLines(inputSlices[1:], parserOut.Grid)
	if err != nil {
		return nil, fmt.Errorf("failed to get vehicles from lines, %v", err)
	}

	parserOut.Vicheles = getVehiclesFromLines

	return &parserOut, nil
}

func getVehiclesFromLines(vehicleStrings [][]string, grid *plateau.Plateau) ([]vehicle.Vehicle, error) {
	out := make([]vehicle.Vehicle, 0)

	vehicleStringsCount := len(vehicleStrings)
	if vehicleStringsCount == 0 {
		return out, fmt.Errorf("no vehicle strings supplied")
	}

	for i := 0; i < vehicleStringsCount; i += 2 {
		vehicleParts := vehicleStrings[i:]
		if len(vehicleParts) < expectedVehicleStringParts {
			return out, fmt.Errorf("incorrectly found vehicle parts, expected two slice indexes")
		}

		startingPositions, err := parseStartingPosition(vehicleParts[0])
		if err != nil {
			return out, err
		}

		rover, err := rover.New(rover.RoverInArgs{
			Position: coordinates.Coordinate{
				X: startingPositions.X,
				Y: startingPositions.Y,
			},
			DirectionString: startingPositions.Direction,
		})
		if err != nil {
			return out, fmt.Errorf("failed to create vehicle, %v", err)
		}

		vichelesCommands := vehicleParts[1]
		vichelesCommandsString := vichelesCommands[0]
		if err := rover.Navigate(vichelesCommandsString, grid); err != nil {
			return out, err
		}

		out = append(out, rover)
	}

	return out, nil
}

func parseStartingPosition(positionSlice []string) (startingPositions, error) {
	if len(positionSlice) == 0 {
		return startingPositions{}, fmt.Errorf("empty vehicle position supplied")
	}

	positionString := positionSlice[0]
	positionStringFormatted := strings.ReplaceAll(positionString, " ", "")

	if len(positionStringFormatted) != expectedStartingPositions {
		return startingPositions{}, fmt.Errorf("invalid vehicle starting position supplied: %s", positionStringFormatted)
	}

	out := startingPositions{}

	startingX, err := strconv.Atoi(string(positionStringFormatted[0]))
	if err != nil {
		return out, fmt.Errorf("failed to parse starting position X: %v", err)
	}

	startingY, err := strconv.Atoi(string(positionStringFormatted[1]))
	if err != nil {
		return out, fmt.Errorf("failed to parse starting position X: %v", err)
	}

	return startingPositions{
		X:         startingX,
		Y:         startingY,
		Direction: string(positionStringFormatted[2]),
	}, nil
}

func validateGridSize(gridSlice []string) (int, int, error) {
	gridStringFormatted := strings.ReplaceAll(gridSlice[0], " ", "")

	if len(gridStringFormatted) != gridFieldsCount {
		return 0, 0, fmt.Errorf("unexpected grid string supplied: %s", gridStringFormatted)
	}

	// Convert string numbers to integers.
	width, err := strconv.Atoi(string(gridStringFormatted[0]))
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read grid width: %v", err)
	}

	height, err := strconv.Atoi(string(gridStringFormatted[1]))
	if err != nil {
		return 0, 0, fmt.Errorf("failed to read height width: %v", err)
	}

	if width == 0 || height == 0 {
		return 0, 0, errors.New("empty grid supplied")
	}

	return width, height, nil
}

// splitLines takes a string of data and returns each input line as a index in a slice.
func splitLines(s string) [][]string {
	var inputLines []string
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		inputLines = append(inputLines, strings.TrimSpace(sc.Text()))
	}

	formattedLines := make([][]string, 0)
	for _, lineString := range inputLines {
		lineValues := make([]string, 0)
		if strings.TrimSpace(lineString) == "" {
			continue
		}

		lineValues = append(lineValues, lineString)
		formattedLines = append(formattedLines, lineValues)
	}

	return formattedLines
}
