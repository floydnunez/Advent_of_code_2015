package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := os.ReadFile("02/real.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	fmt.Println(lines)
	total := 0
	total_ribbon := 0
	for _, str := range lines {
		dim := strings.Split(str, "x")
		l, _ := strconv.Atoi(dim[0])
		w, _ := strconv.Atoi(dim[1])
		h, _ := strconv.Atoi(dim[2])
		area := 2*l*w + 2*w*h + 2*h*l
		area += min(l*w, w*h, h*l)
		total += area

		total_ribbon += 2 * min(l+w, w+h, l+h)
		total_ribbon += l * w * h
	}
	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", total_ribbon)

}

func min(l int, w int, h int) int {
	return min2(l, min2(w, h))
}

func min2(w int, h int) int {
	if w < h {
		return w
	}
	return h
}
