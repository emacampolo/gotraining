package main

import (
	`fmt`
)

type alice struct {
	name string
}

type bob struct {
	name string
}

func main() {
	// Declare alice and bob and sets to its zero value
	var a alice
	var b bob

	fmt.Println(a, b)
	// a = b // will implicit conversion work ? and explicit ?

	c := struct {
		name string
	}{
		name: "name",
	}

	fmt.Println(c)

	// Assign the value of the unnamed struct type
	// to the named struct type value.
	// a = c
}
