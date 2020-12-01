package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	ids := []int{33, 2, 76, 23, 11, 4}

	// new for loop, using i as an index over the collection of ids
	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// or if you don't need the index use a _
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	// use range on a map
	emails2 := map[string]string{
		"Bob": "bob@gmail.com",
		"Joe": "joe@gmail.com",
		"Sue": "sue@gmail.com",
	}

	for k, v := range emails2 {
		fmt.Printf("%s: %s\n", k, v)
	}
}
