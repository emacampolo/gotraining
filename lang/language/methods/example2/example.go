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
	// A function is a 2 word structure.
	// The first word points to the code for that method we want to execute.
	// The second word is a pointer to the copy of data (displayName uses a value receiver)
	//  -----
	// |  *  | --> code
	//  -----
	// |  *  | --> copy of bill
	//  -----
	bill.name = "Bill"
	f1()
	//  -----
	// |  *  | --> code
	//  -----
	// |  *  | --> bill
	//  -----
	f2 := bill.setName

	// Call the method via the variable.
	f2("Bill")
	bill.displayName()
}
// END OMIT
