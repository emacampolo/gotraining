package main

import "fmt"

func main() {
	// Using the value semantic form of the for range.
	friends := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for _, v := range friends {
		friends = []string{}
		fmt.Println(v)
	}
}
