// Sample program to show the basic concept of pass by value.
package main

import (
	`fmt`
)

func main() {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	fmt.Println("count: Value of", count, " Addr of", &count)

	// Pass the "value of" the count. We slice a new frame off the stack
	// Every time we call a function we cross a program boundary.
	// It's creating a sandbox, a layer of isolation,
	// What's going to happen is we're going to slice a new frame off the stack.
	increment(count)

	fmt.Println("count: Value of", count, " Addr of", &count)
}

// We get a copy of the value of count. We are crossing a program boundary. Inc will need 8 bytes in the new frame.
func increment(inc int) {

	// Read, modify and write operation. We increment the "value of" inc.
	// This is an isolated modification. No side effect.
	inc++
	fmt.Println("count: Value of", inc, " Addr of", &inc)
}
