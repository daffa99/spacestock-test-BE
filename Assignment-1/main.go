package main

import (
	"fmt"
)

func sum(numArray []int) int {
	// Take a number and remove it from numArray
	number := numArray[0]
	newArray := numArray[1:]
	// If that number is the last number
	if len(newArray) == 0 {
		return number
	}
	// Sum the number with the rest of array
	return number + sum(newArray)
}

func main() {
	// Test Case
	numArray := []int{
		1,
		2,
		5,
		10,
	}
	fmt.Println(sum(numArray))
}
