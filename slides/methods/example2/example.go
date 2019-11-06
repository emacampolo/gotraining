package main

import "fmt"

type user struct {
	name string
}

func (d user) displayName() {
	fmt.Println("My Name Is", d.name)
}

func (d *user) setName(name string) {
	d.name = name
}

// START OMIT
func main() {
	bill := user{}

	f1 := bill.displayName
	f1()
	bill.name = "Bill"
	f1()

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get the address of d because the method
	// is using a pointer receiver.
	f2 := bill.setName

	// Call the method via the variable.
	f2("Bill")

	bill.displayName()

	// Change the value of d.
	bill.name = "Edu"

	// Call the method via the variable. We see the change.
	bill.displayName()
}
// END OMIT
