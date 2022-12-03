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

type pos struct {
	x int
	y int
}

func main() {
	dat, err := os.ReadFile("03/real.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	fmt.Println(lines)
	fmt.Println(lines[0])
	{
		current := pos{0, 0}
		visited := make(map[pos]int)
		visited[current] = 1
		for _, mov := range lines[0] {
			current = move(mov, current)
			if _, ok := visited[current]; ok {
				visited[current] += 1
			} else {
				visited[current] = 1
			}
		}
		fmt.Println("Part 1:", len(visited))
	}
	{
		santa := pos{0, 0}
		robot := pos{0, 0}
		visited := make(map[pos]int)
		visited[santa] = 2
		for turn, mov := range lines[0] {
			var current pos
			if turn%2 == 0 {
				current = santa
			} else {
				current = robot
			}
			current = move(mov, current)
			if _, ok := visited[current]; ok {
				visited[current] += 1
			} else {
				visited[current] = 1
			}
			if turn%2 == 0 {
				santa = current
			} else {
				robot = current
			}
		}
		fmt.Println("Part 2:", len(visited))
	}
}

func move(mov int32, current pos) pos {
	if mov == '^' {
		current.y += 1
	} else if mov == '>' {
		current.x += 1
	} else if mov == '<' {
		current.x -= 1
	} else if mov == 'v' {
		current.y -= 1
	}
	return current
}
