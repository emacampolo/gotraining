package main

import "fmt"

// START OMIT
func main() {
	count := 10

	fmt.Println("count: Value of", count, " Addr of", &count)

	// Pass the "value of" the count. We slice a new frame off the stack
	increment(count)

	fmt.Println("count: Value of", count, " Addr of", &count)
}

// We get a copy of the value of count. We are crossing a program boundary.
// Inc will need 8 bytes in the new frame.
func increment(inc int) {

	// Read, modify and write operation. We increment the "value of" inc.
	// This is an isolated modification. No side effect.
	inc++
	fmt.Println("count: Value of", inc, " Addr of", &inc)
}
// END OMIT