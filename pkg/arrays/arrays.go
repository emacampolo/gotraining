// Sample program to show how to declare and iterate over
// arrays of different types.
package main

import "fmt"

func main() {

	// Declare an array of five strings that is initialized
	// to its zero value.
	var fruits [5]string
	fruits[0] = "Apple" // creates a string and assign a copy
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// We use the value semantics. fruit variable is another copy.
	for i, fruit := range fruits {
		fmt.Println(i, fruit) // share down the stack another copy.
	}
	// only one of each string is allocated in the heap.
}
