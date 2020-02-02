package main

import "fmt"

func main() {
	friends := [5]string{"Annie", "Betty", "Charley", "Doug", "Edward"}

	for i := range friends {
		friends[1] = "Jack"

		if i == 1 {
			fmt.Println(friends[1])
		}
	}
}
