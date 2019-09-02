// Sample program to show how the capacity of the slice
// is not available for use.
package main

import "fmt"

func main() {

	// Create a slice with a length of 5 elements. Make is a built-in types for create reference types.
	// A slice is a 3-word data structure
	// [*, len, cap]

	// pointer: points to a backing array
	// len: total elements I can access from the pointer position
	// cap: total elements that exists from the pointer position

	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// You can't access an index of a slice beyond its length.
	//fruits[5] = "Runtime error"

	// Passing a value to print function. Slices are designed to be copied and leverage value semantics.
	// We can pass thousands of slices. Only the array will be allocated.
	fmt.Println(fruits)
}