package main

import "fmt"

// START OMIT
func main() {

	var i int
	var f float64
	var b bool

	// Strings are two words data structure (8 or 16 bytes)
	// type StringHeader struct {
	//     Data uintptr
	//     Len  int
	// }
	var s string

	fmt.Printf("var i int \t %T [%v]\n", i, i)
	fmt.Printf("var f float64 \t %T [%v]\n", f, f)
	fmt.Printf("var b bool \t %T [%v]\n\n", b, b)
	fmt.Printf("var s string \t %T [%v]\n", s, s)
}
// END OMIT
