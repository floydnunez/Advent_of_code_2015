package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type instruction struct {
	kind int
	inix int
	iniy int
	finx int
	finy int
}

func main() {
	fileLines, err := readFileIntoArray("06/real.txt")
	if err {
		println("Error fatal")
		return
	}

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
	turnLightsInRectangle(0, grid, 1, 1, 500, 100)
	numLightsOn := countLightsOn(grid)
	fmt.Println("Number of lights on:", numLightsOn)

	var instructions []instruction
	naughtyRegex := regexp.MustCompile(`(.{1,8}) (\d+),(\d+) (.{7}) (\d+),(\d+)`)
	for _, line := range fileLines {
		matches := naughtyRegex.FindAllStringSubmatch(line, -1)
		for _, elem := range matches {
			if len(elem) != 7 {
				panic("aaaah")
			}
			var kind int
			if elem[1] == "toggle" {
				kind = 2
			} else if elem[1] == "turn on" {
				kind = 1
			} else if elem[1] == "turn off" {
				kind = 0
			} else {
				panic("eeeeeh?")
			}
			inix, err1x := strconv.Atoi(elem[2])
			iniy, err1y := strconv.Atoi(elem[3])
			finx, err2x := strconv.Atoi(elem[5])
			finy, err2y := strconv.Atoi(elem[6])
			if err1x != nil || err1y != nil || err2x != nil || err2y != nil {
				fmt.Println(err1x, err1y, err2x, err2y)
				panic("iiiiiih? ")
			}
			instructions = append(instructions, instruction{kind: kind, inix: inix, iniy: iniy, finx: finx, finy: finy})
			break
		}
	}
	for _, inst := range instructions {
		fmt.Println("inst:", inst)
		turnLightsInRectangle(inst.kind, grid, inst.inix, inst.iniy, inst.finx, inst.finy)
	}
	numLightsOn = countLightsOn(grid)
	fmt.Println("Part 1: Number of lights on:", numLightsOn)

	turnLightsInRectangle(0, grid, 0, 0, 999, 999)
	for _, inst := range instructions {
		fmt.Println("inst:", inst)
		adjustBrightnessLightsInRectangle(inst.kind, grid, inst.inix, inst.iniy, inst.finx, inst.finy)
	}
	totalBrightness := countBrightness(grid)
	fmt.Println("Part 2: total brightness:", totalBrightness)

}

func countBrightness(grid [][]int) interface{} {
	totalBrightness := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			totalBrightness += grid[i][j]
		}
	}
	return totalBrightness
}

func readFileIntoArray(filename string) ([]string, bool) {
	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil, true
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	errclose := readFile.Close()
	return fileLines, errclose != nil
}

func countLightsOn(grid [][]int) int {
	numLightsOn := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] == 1 {
				numLightsOn++
			}
		}
	}
	return numLightsOn
}

func turnLightsInRectangle(val int, grid [][]int, x1 int, y1 int, x2 int, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			newval := val
			if val < 0 || val > 1 {
				newval = (grid[i][j] + 1) % 2
			}
			grid[i][j] = newval
		}
	}
}

func adjustBrightnessLightsInRectangle(val int, grid [][]int, x1 int, y1 int, x2 int, y2 int) {
	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			if val == 0 && grid[i][j] > 0 {
				grid[i][j] = grid[i][j] - 1
			}
			if val == 1 {
				grid[i][j] = grid[i][j] + 1
			}
			if val == 2 {
				grid[i][j] = grid[i][j] + 2
			}
		}
	}
}
