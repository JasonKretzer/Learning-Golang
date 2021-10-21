package main

import (
	"fmt"
	"strconv"
)

// pretty much looks like C

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

// this adds a function to the struct above (this is a value receiver function)
// it takes the Person struct by value
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old."
}

// note this takes a pointer so you can change the underlying values
func (p *Person) haBirthday() {
	p.age++
}

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}

}

func main() {
	// you can remove the keys below, too
	person1 := Person{firstName: "Sam", lastName: "Smith", city: "Boston", gender: "f", age: 25}
	fmt.Println(person1)
	fmt.Println(person1.greet())
	person1.haBirthday()
	person1.getMarried("Jones")
	fmt.Println(person1.greet())
}
