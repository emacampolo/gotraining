package main

import "fmt"

func main() {
	friends := [5]string{"1", "2", "3", "4", "5"}

	for i, v := range friends {
		friends[1] = "9"

		if i == 1 {
			fmt.Println(v)
		}
	}
}
