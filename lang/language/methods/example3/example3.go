package main

import "fmt"

// START OMIT
type user struct {
	name  string
	likes int
}

func (u *user) notify() {
	fmt.Printf("%s has %d likes\n", u.name, u.likes)
}

func (u *user) addLike() {
	u.likes++
}

func main() {
	users := []user{{name: "bill"}, {name: "lisa"}}

	for _, u := range users {
		u.addLike()
	}

	for _, u := range users {
		u.notify()
	}
}

//END OMIT
