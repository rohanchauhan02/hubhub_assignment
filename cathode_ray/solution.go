package cathoderay

import (
	"fmt"
	"strconv"
	"strings"
)

func CalculateSignalStrength(instructions []string) int {
	X := 1
	signalStrength := 0
	cycle := 0

	// iterate over instructions
	for _, instruction := range instructions {
		cycle++
		parts := strings.Split(instruction, " ")
		operation := parts[0]
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			// fmt.Printf("Cycle1: %v, X: %v \n", cycle, X)
			signalStrength += cycle * X
		}
		if operation == "addx" {
			value, err := strconv.Atoi(parts[1])
			if err != nil {
				fmt.Printf("Failed to convert string to integer, String: %s", parts[1])
			}
			X += value
			cycle++
			if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
				// fmt.Printf("Cycle2: %v, X: %v \n", cycle, X)
				signalStrength += cycle * X
			}
		}

	}
	return signalStrength
}
