// Sample program to show how the for range has both value and pointer semantics.
package main

import "fmt"

func main() {

	// Using the pointer semantic form of the for range.
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("[%s] : ", friends[1])

	for i := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("[%s]\n", friends[1])
		}
	}

/*
	// Using the value semantic form of the for range.
	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("[%s] : ", friends[1])

	for i, v := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("[%s]\n", v)
		}
	}
*/

/*
	// Using the value semantic form of the for range but with pointer
	// semantic access. DON'T DO THIS.
	friends = [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	fmt.Printf("[%s] : ", friends[1])

	// copy of the address of the array
	for i, v := range &friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Printf("[%s]\n", v)
		}
	}
*/
}
