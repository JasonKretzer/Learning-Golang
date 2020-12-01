package main

import "fmt"

func main() {

	// make provides memory allocation --
	// looks like stack allocation as it does not create a pointer
	// there is also a "new" keyword for allocation that does return a pointer

	// maps are standard key/value pair stores

	// map syntax => map[KeyType]ValueType
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Joe"] = "joe@gmail.com"
	emails["Sue"] = "sue@gmail.com"

	fmt.Println(emails)

	// to delete an entry
	delete(emails, "Bob")

	// can also declare and assign

	emails2 := map[string]string{
		"Bob": "bob@gmail.com",
		"Joe": "joe@gmail.com",
		"Sue": "sue@gmail.com",
	}

	fmt.Println(emails)
	fmt.Println(emails2)

}
