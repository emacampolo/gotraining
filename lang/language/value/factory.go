package main

type user struct {
	name string
}

func main() {
	// We defined factory functions. We don't call it constructors ;)
	// They tell use which semantic are in play.
	u := createUser()
	println("u in main", &u)
}

// createUser creates a user value and passed a copy back to the caller.
func createUser() user {
	u := user{name: "Gopher"}
	println("u in function", &u)
	// We return a copy of u
	return u
}
