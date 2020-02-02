package main

type user struct {
	name string
}

func main() {
	// We defined factory functions. We don't call it constructors ;)
	// They tell use which semantic are in play.
	u := createUser()
	println("u in main", u)
}

// createUser creates a user value and shares the value with the caller.
func createUser() *user {
	// Construction tell us nothing if the value will stay on the stack or heap.
	// Only how the value it is shared.
	u := user{name: "Gopher"}
	println("u in function", &u)
	// Allocation is needed. & is used as a readability marker.
	return &u
}
