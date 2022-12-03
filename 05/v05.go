package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Open the input file
	file, err := os.Open("05/example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a scanner to read the input file line by line
	scanner := bufio.NewScanner(file)

	// Regular expression to match three or more vowels
	vowelRegex := regexp.MustCompile(`[aeiou]{3,}`)

	// Regular expression to match two or more identical letters
	doubleLetterRegex := regexp.MustCompile(`([a-zA-Z])\\1+`)

	// Regular expression to match the disallowed substrings
	naughtyRegex := regexp.MustCompile(`ab|cd|pq|xy`)

	// Counter for the number of nice strings
	niceStrings := 0

	// Read the input file line by line
	for scanner.Scan() {
		line := scanner.Text()
		println(vowelRegex.MatchString(line), doubleLetterRegex.MatchString(line), !naughtyRegex.MatchString(line))
		// Check if the line contains at least three vowels, a double letter, and no disallowed substrings
		if vowelRegex.MatchString(line) && doubleLetterRegex.MatchString(line) && !naughtyRegex.MatchString(line) {
			// If all conditions are met, increment the nice strings counter
			niceStrings++
		}
	}

	// Print the number of nice strings
	fmt.Println(niceStrings)
}
