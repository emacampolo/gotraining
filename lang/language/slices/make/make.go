package main

import (
	"fmt"
)

// START OMIT
func main() {
	s := make([]int, 5)
	// s[1] = 1

	describe(s)

	s = make([]int, 0, 5)

	// s[1] = 1
	describe(s)
}
// END OMIT

func describe(s []int) {
	fmt.Printf("slice: %v | len: %d | cap: %d\n", s, len(s), cap(s))
}
