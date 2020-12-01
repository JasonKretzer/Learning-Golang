package main

import "fmt"

// function called adder that returns a function that takes an int and returns an int
//   name    ->return type
func adder() func(int) int {
	sum := 0
	// function to return
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {

	sum := adder()

	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}

}
