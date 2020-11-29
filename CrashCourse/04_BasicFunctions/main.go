package main

// how to import multiple packages
import (
	"fmt"
)

//basic function signature formatting is:
//func nameOfFunction(paramName paramType) returnType
func greeting(name string) string {
	return "Hello, " + name
}

//can use a list for parameters like below when the same type
func intSum(n1, n2 int) int {
	return n1 + n2
}

func main() {
	fmt.Println(greeting("Penelope"))
	fmt.Println(intSum(1, 3))
}
