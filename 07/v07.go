package main

import (
	"fmt"
)

func main() {
	// Create a grid of 1000 by 1000 lights
	grid := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		grid[i] = make([]int, 1000)
	}

	// Initialize all the lights to be off
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			grid[i][j] = 0
		}
	}
	turnOnLightsInRectangle(grid, 0, 0, 999, 999)
	// Print the number of lights that are on
	numLightsOn := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] == 1 {
				numLightsOn++
			}
		}
	}
	fmt.Println("Number of lights on:", numLightsOn)
}

func turnOnLightsInRectangle(grid [][]int, x1 int, y1 int, x2 int, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grid[i][j] = 1
		}
	}
}
