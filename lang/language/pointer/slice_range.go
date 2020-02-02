package main

import "fmt"

func main() {
	friends := []string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for i := range friends {
		friends = []string{}
		fmt.Println(friends[i])
	}
}
