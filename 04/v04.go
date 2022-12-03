package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

func main() {
	// The secret key
	secretKey := "bgvyzdsv"

	// Iterate over a range of positive numbers
	for i := 1; i < 10000000; i++ {
		// Compute the MD5 hash of the secret key concatenated with the current number
		md5Hash := fmt.Sprintf("%x", md5.Sum([]byte(secretKey+strconv.Itoa(i))))
		// Check if the hash starts with six zeroes
		if md5Hash[:5] == "00000" {
			// If it does, print the number and break the loop
			fmt.Println("Part 1:", i)
			break
		}
	}
	// Iterate over a range of positive numbers
	for i := 1; i < 10000000; i++ {
		// Compute the MD5 hash of the secret key concatenated with the current number
		md5Hash := fmt.Sprintf("%x", md5.Sum([]byte(secretKey+strconv.Itoa(i))))
		// Check if the hash starts with six zeroes
		if md5Hash[:6] == "000000" {
			// If it does, print the number and break the loop
			fmt.Println("Part 2:", i)
			break
		}
	}

}
