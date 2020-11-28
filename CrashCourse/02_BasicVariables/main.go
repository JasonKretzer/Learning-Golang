package main

import "fmt"

func main() {

	//if you declare a variable you have to use it or it is an error
	// you can use the string or not here as it can infer the type
	var name string = "Jason"
	var name2 = "Jason"
	var age = 45

	//const -- cannot assign to it after created
	const isCool = true

	//also variable declaration shorthand
	name3 := "Noname"
	size := 1.342321321

	//finally multi variable declaration/assignments
	name4, name5 := "Polyphemus", "Odysseus"

	fmt.Println(name, name2, isCool, name3, name4, name5, size)

	// go has format specifiers in printf similar to C
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
}
