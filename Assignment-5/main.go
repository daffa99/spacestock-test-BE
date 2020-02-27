package main

import "fmt"

//Case struct
type Case struct {
	Order int
	Name  string
}

var cases = []Case{
	{Order: 1, Name: "one"},
	{Order: 2, Name: "two"},
	{Order: 3, Name: "tri"},
	{Order: 4, Name: "fou"},
	{Order: 5, Name: "fiv"},
}

// Rearrange func. Conditions: where the difference between 2 args is 2 (skipped 1), it will rearrange one another AND the biggest one and the one in between, while the other will just rearrange between each other
func Rearrange(object []Case, before int, after int) []Case {
	if before > after {
		if before-after == 2 {
			nameBefore := object[before-1].Name
			nameAfter := object[after-1].Name
			nameMiddle := object[before-2].Name

			object[before-1].Name = nameMiddle
			object[before-2].Name = nameAfter
			object[after-1].Name = nameBefore
		} else {
			nameBefore := object[before-1].Name
			nameAfter := object[after-1].Name

			object[before-1].Name = nameAfter
			object[after-1].Name = nameBefore
		}
	} else {
		if after-before == 2 {
			nameAfter := object[after-1].Name
			nameBefore := object[before-1].Name
			nameMiddle := object[after-2].Name

			object[after-1].Name = nameMiddle
			object[after-2].Name = nameBefore
			object[after-1].Name = nameAfter
		} else {
			nameBefore := object[before-1].Name
			nameAfter := object[after-1].Name

			object[before-1].Name = nameAfter
			object[after-1].Name = nameBefore
		}
	}
	return object
}

func main() {
	// Test case 1
	cases = Rearrange(cases, 1, 2)
	cases = Rearrange(cases, 3, 1)
	fmt.Printf("%+v\n", cases)
	/*
	  Expect cases = [
	    {Order: 1, Name: "tri"},
	    {Order: 2, Name: "two"},
	    {Order: 3, Name: "one"},
	    {Order: 4, Name: "fou"},
	    {Order: 5, Name: "fiv"},
	  ] */

	// Test case 2
	// cases = Rearrange(cases, 4, 2)
	// cases = Rearrange(cases, 1, 2)
	// fmt.Printf("%+v\n", cases)
	/*
	  Expect cases = [
	    {Order: 1, Name: "fou"},
	    {Order: 2, Name: "one"},
	    {Order: 3, Name: "two"},
	    {Order: 4, Name: "tri"},
	    {Order: 5, Name: "fiv"},
	  ]
	*/
}
