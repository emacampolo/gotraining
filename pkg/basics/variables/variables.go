package main

import (
	`fmt`
)

func main() {
	// We specify type and size
	var c float64

	// Declare variables that are set to their zero value. We only know that is an int.
	// Size will depend on the architecture. Address size can be 64 (8 bytes) or 32 bits (4 bytes).
	var a int

	// Don't explicit declare size if it is not obvious
	// var x int64

	// Strings are two words data structure
	// On the Playground the word is gonna be two, four byte words,
	// so that's eight bytes and on our 64 bit, it'll be 16 two eight byte words
	// [ nil , 0 ]
	var b string

	// 1 byte, either true or false
	var d bool

	// +**
	// Even if there are multiple ways to declare a variable, var is use as a readability marker
	// ***

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 10

	// [ ptr -> []byte{"hello"} | 5 ]
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// nothing prevents us from doing xx := 0 but we should use it consistently.
	// Prefer var for zero value semantics instead.
}
