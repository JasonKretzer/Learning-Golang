package main

import "fmt"

// how to import multiple packages

func main() {

	// ***********************************
	// ARRAYS
	// ***********************************

	// note that the [] are part of the type and are next to it
	var fruits [2]string

	//assign like C based languages do
	fruits[0] = "Apple"
	fruits[1] = "Orange"

	fmt.Println(fruits)
	fmt.Println(fruits[0])
	fmt.Println(fruits[1])

	//can also declare and assign at the same time
	veggies := [2]string{"Lettuce", "Carrots"}

	fmt.Println(veggies)

	// ***********************************
	// SLICES -- just like array but do not need length up front
	// ***********************************

	fruits2 := []string{"Banana", "Grape", "Plum"}

	fmt.Println(fruits2)

	//len gets the length
	fmt.Println(len(fruits2))

	//get a range of items from the array or slice
	//note the range values are indexes and is left inclusive but not right inclusive
	fmt.Println(fruits2[1:1])

	//you can append to the end of slices as well
	fruits2 = append(fruits2, "Watermelon")
	fmt.Println(fruits2)
}
