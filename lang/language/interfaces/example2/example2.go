package main

import "fmt"

// START OMIT
type greeter interface {
	greet()
}

type user struct {
	name string
}

func (u *user) greet() {
	fmt.Printf("Hi, My name is: %s", u.name)
}

func main() {
	u := user{"Bill"}

	// compiler prevents you from storing copies of values (value semantics)
	// inside the interface if you implement the interface using pointer semantics
	// You can never assume that it is safe to make a copy of any value that is pointed to by a pointer.
	sendGreetings(u)
}

func sendGreetings(g greeter) {
	g.greet()
}

// END OMIT
