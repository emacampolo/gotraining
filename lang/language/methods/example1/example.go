package main

import "fmt"

// START OMIT
type user struct {
	name string
}

func (d user) displayName() {
	fmt.Println("My Name Is", d.name)
}

func (d *user) setName(name string) {
	d.name = name
}

func main() {
	bill := user{}

	bill.setName("Bill")
	bill.displayName()

	edu := &user{}
	edu.setName("Edu")
	edu.displayName()
}
// END OMIT
