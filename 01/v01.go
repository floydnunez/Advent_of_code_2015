package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("01/real.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	fmt.Println("lines[0]", lines[0])
	p1_floor := 0
	position := 0
	for number, letter := range lines[0] {
		if letter == '(' {
			p1_floor += 1
		}
		if letter == ')' {
			p1_floor -= 1
		}
		fmt.Println(p1_floor+1, position)
		if p1_floor == -1 && position == 0 {
			position = number + 1
		}
	}
	fmt.Println("Part 1: ", p1_floor) //232
	fmt.Println("Part 2: ", position) //232

}
