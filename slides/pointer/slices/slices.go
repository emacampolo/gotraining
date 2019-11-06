package main

import "fmt"

func main() {
	// Using the pointer semantic form of the for range.
	friends := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for i := range friends {
		friends = []string{}
		fmt.Println(friends[i])
	}
}