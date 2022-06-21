# Mars Rover Challenge

## Running the project

`main.go` contains a hardcoded string input which directly calls the `parser.Parse()` function, returning the vehicles + grid size. 

This could be future extended using the Go standard lib `flag` package to accept instructions into the program.

To see the result of the hardcoded input please run:

`go run main.go`.


## Running tests

`$ ./testandlint.sh`

Running `testandlint.sh` requires the `golangci-lint` Go linter installed on your machine.

Install `golangci-lint`

- [Installation](https://golangci-lint.run/usage/install/#local-installation)
- [Github](https://github.com/golangci/golangci-lint)


## Packages
-   `cooordinates` Handles coordinate movements on vehicles, and supplies an interface + mocks.
- `direction` Handles rotating the vehicle.
- `parser` Handles string input for vehicle instructions and plateau size. It also directly calls the vehicle nagivation function to execute the instructions.
- `vehicle` supplies an interface with the required implementation for a vehicle. `vehicle/rover` satisfies this interface.
- `plateau` initialises a new mars rover grid with the supplied grid width and grid height from the string instructions.

## Example input

```
5 5

1 2 N

LMLMLMLMM

3 3 E

MMRMMRMRRM
```

## Expected output of example

```
1 3 N

5 1 E
```

## About the challenge

A squad of robotic rovers are to be landed by NASA on a plateau on Mars.
This plateau, which is curiously rectangular, must be navigated by the rovers so that their on board cameras can get a complete view of the surrounding terrain to send back to Earth.
A rover's position is represented by a combination of an x and y co-ordinates and a letter representing one of the four cardinal compass points. The plateau is divided up into a grid to simplify navigation. An example position might be 0, 0, N, which means the rover is in the bottom left corner and facing North.
In order to control a rover, NASA sends a simple string of letters. The possible letters are 'L', 'R' and 'M'. 'L' and 'R' makes the rover spin 90 degrees left or right respectively, without moving from its current spot.
'M' means move forward one grid point, and maintain the same heading. Assume that the square directly North from (x, y) is (x, y+1).