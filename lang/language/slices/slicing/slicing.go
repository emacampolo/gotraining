package main

import (
	"fmt"
)

// START OMIT
func main() {
	s1 := []int{0, 1, 2, 3, 4, 5}
	describe(s1)

	s2 := s1[1:]
	describe(s2)

	s3 := s1[:3]
	describe(s3)

	s4 := s1[2:4]
	describe(s4)

	s4[0] = 99
	describe(s1)
}

// END OMIT

func describe(s []int) {
	fmt.Printf("slice: %v | len: %d | cap: %d\n", s, len(s), cap(s))
}
