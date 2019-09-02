package main

import "fmt"

func main() {

	// Declare an array of 5 strings initialized with values.
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	// Iterate over the array displaying the value and
	// address of each element.
	// We can see the continuous block of addresses. Since we are in amd64,
	// each string is a 2-word data structure of 8 bytes each.
	for i, v := range friends {
		fmt.Printf("Value[%s]\tAddress[%p] IndexAddr[%p]\n", v, &v, &friends[i])
	}
}