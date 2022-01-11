package main

import "fmt"

type deck []string

// function that has a receiver inside the parens before the function name
// basically, what ever is calling this function is passed into the function itself
// kind of like a 'this'
// also, it is a convention to use the first letter as the variable name for the receiver object
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

/* for this function the d is a working copy of the deck variable that
* was used to call the function.  It is then available in the function
* when you reference d.  Every variable of type deck has access to this function
 */
