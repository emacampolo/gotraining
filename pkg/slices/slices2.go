// Sample program to show how to grow a slice using the built-in function append
// and how append grows the capacity of the underlying array.
package main

import "fmt"

func main() {

	// Declare a nil slice of strings.
	// We are using the zero value of a reference type (the empty literal []string{} is not the zero value)->[* ptr -> struct{}, 0, 0]
	var data []string
	// data := make([]string, 0, 1e5) we never have to make a copy of the backing array.

	// Capture the capacity of the slice.
	lastCap := cap(data)

	// Append ~100k strings to the slice.
	for record := 1; record <= 1e5; record++ {

		// Use the built-in function append to add to the slice.
		value := fmt.Sprintf("Rec: %d", record)

		// Append uses value semantics.
		data = append(data, value)

		// When the capacity of the slice changes, display the changes.
		if lastCap != cap(data) {

			// Calculate the percent of change.
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// Save the new values for capacity.
			lastCap = cap(data)

			// Display the results.
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0],
				record,
				cap(data),
				capChg)
		}
	}
}