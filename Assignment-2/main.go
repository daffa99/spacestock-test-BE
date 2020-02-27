package main

import (
	"fmt"
)

func contains(strArray []string, check string) bool {
	word := strArray[0]
	// Check the word
	if word == check {
		return true
	}
	// Remove checked word in strArray
	newArray := strArray[1:]
	// If no other word to check
	if len(newArray) == 0 {
		return false
	}
	return contains(newArray, check)
}

func main() {
	// Test case
	strArray := []string{
		"Makan",
		"Tidur",
		"Mandi",
		"Sholat",
		"Main",
	}
	fmt.Println(contains(strArray, "Belajar"))
	fmt.Println(contains(strArray, "Main"))
}
