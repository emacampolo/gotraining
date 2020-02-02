package main

import "fmt"

// START OMIT
func main() {
	count := 10

	fmt.Printf("Value     :%d\n", count)
	fmt.Printf("Addr      :%p\n\n", &count)

	// Pass the "address of" count. We still pass by value but it happens to be an address.
	// This is not pass by reference!
	increment(&count)

	fmt.Printf("Value     :%d\n", count)
	fmt.Printf("Addr      :%p\n\n", &count)
}

// increment declares inc as a pointer variable whose value is an address
// and points to values of type int.
func increment(inc *int) {
	// Indirect read, modify, write operation.
	// We are having side effects since we no longer acts in a sandbox.
	*inc++
	fmt.Printf("Value     :%d\n", inc)
	fmt.Printf("Addr      :%p\n", &inc)
	fmt.Printf("Points to :%d\n\n", *inc)
}
// END OMIT
