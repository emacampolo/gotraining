package main

import "fmt"

// START OMIT
func main() {
	count := 10

	fmt.Println("count: Value of", count, "Addr of", &count)

	// Pass the "address of" count. We still pass by value but it happens to be an address.
	// This is not pass by reference!
	increment(&count)

	fmt.Println("count: Value of", count, "Addr of", &count)
}

// increment declares inc as a pointer variable whose value is an address
// and points to values of type int.
func increment(inc *int) {

	// Indirect read, modify, write operation.
	// We are having side effects since we no longer acts in a sandbox.
	*inc++
	fmt.Println("inc: Value of", inc, "Addr of", &inc, "Value Points To", *inc)
}
// END OMIT
