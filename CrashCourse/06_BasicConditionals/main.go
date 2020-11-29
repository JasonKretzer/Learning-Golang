package main

import "fmt"

func main() {
	x := 5
	y := 10

	// typically don't need() but can use if you want
	if x < y {
		fmt.Println("x is less than y")
	} else { //has standard else and if else constructs as well
		fmt.Println("y is less than x")
	}
	//can also do:
	// ==, <=, => for equality, lte, and gte
	// &&, || logical and/or

	// switch
	color := "red"

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is not blue or red")
	}
}
