package main

import "fmt"

type greeter interface {
	greet()
}

type user struct {
	name string
}

func (u *user) greet() {
	fmt.Printf("Hi, My name is: %s", u.name)
}

func main() {
	u := user{"Bill"}
	sendGreetings(u)
}

func sendGreetings(g greeter) {
	g.greet()
}
