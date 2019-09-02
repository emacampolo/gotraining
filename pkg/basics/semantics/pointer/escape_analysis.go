// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Sample program to teach the mechanics of escape analysis.
package main

// user represents a user in the system.
type user struct {
	name  string
	email string
}

func main() {

	// We defined factory functions. We don't call it constructors ;). They tell use which semantic are in play.
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", u2)
}

// createUserV1 creates a user value and passed a copy back to the caller. We use value semantics.
//go:noinline
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	// We share Bill to the print function
	println("V1", &u)

	// We return a copy of u
	return u
}

// createUserV2 creates a user value and shares the value with the caller. Pointer semantics.
//go:noinline
func createUserV2() *user {

	// Construction tell us nothing if the value will stay on the stack or heap. Only how the value it is shared.
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u)

	// Allocation is needed. & is used as a readability marker.
	return &u
}

/*
See escape analysis and inlining decisions.
go build -gcflags -m=2 escape_analysis.go
*/