package main

import (
	"fmt"
	"strings"

	bolingboulders "github.com/rohanchauhan02/huhhub_assignment/boling_boulders"
	cathoderay "github.com/rohanchauhan02/huhhub_assignment/cathode_ray"
)

func main() {
	cathode_ray_solution()
	boiling_boulders_solution()
}

func cathode_ray_solution() {
	input := `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`
	instructions := strings.Split(input, "\n")
	signalStrength := cathoderay.CalculateSignalStrength(instructions)
	fmt.Printf("Signal strenth: %v \n", signalStrength)
}

func boiling_boulders_solution() {
	input := `2,2,2
1,2,2
3,2,2
2,1,2
2,3,2
2,2,1
2,2,3
2,2,4
2,2,6
1,2,5
3,2,5
2,1,5
2,3,5`

	var cubes []bolingboulders.Cube
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var cube bolingboulders.Cube
		fmt.Sscanf(line, "%d,%d,%d", &cube.X, &cube.Y, &cube.Z)
		cubes = append(cubes, cube)
	}
	surfaceArea := bolingboulders.CalculateSurfaceArea(cubes)
	fmt.Printf("Surface Area: %d \n", surfaceArea)

}
