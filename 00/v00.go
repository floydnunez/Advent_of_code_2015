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
	dat, err := os.ReadFile("02/example.txt")
	check(err)
	lines := strings.Split(string(dat), "\n")
	fmt.Println(lines)
}
