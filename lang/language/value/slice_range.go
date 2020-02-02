package main

import "fmt"

func main() {
	friends := []string{"1", "2", "3", "4", "5"}

	for _, v := range friends {
		friends = []string{}
		fmt.Print(v)
	}

	/*
	friends = []string{"1", "2", "3", "4", "5"}

	for _, v := range friends {
		friends[3] = "-"
		fmt.Print(v)
	}
	*/
}
