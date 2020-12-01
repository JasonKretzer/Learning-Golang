package main

import "fmt"

func main() {

	// has pointers like C/C++
	// unless you use pointers, everything is passed by value
	a := 5
	b := &a // assigns be the address of a

	fmt.Println(a, b)

	fmt.Printf("a is type %T\n", a) // int
	fmt.Printf("b is type %T\n", b) // *int
	fmt.Printf("*b = %d\n", *b)     // de-reference the pointer -- returns 5

	// can change the value of a by using the *b
	*b = 10
	fmt.Printf("a = %d\n", a)   // now is 10
	fmt.Printf("*b = %d\n", *b) // now is 10

}
