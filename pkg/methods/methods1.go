// Sample program to show how to declare methods and how the Go
// compiler supports them.
package main

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver. It will operates with its own copy of the data.
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver. It wil share data access.
func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {

	// Values of type user can be used to call methods
	// declared with both value and pointer receivers.
	bill := user{"Bill", "bill@email.com"}

	// Why does this compile ?
	bill.changeEmail("bill@hotmail.com")						 // (&bill).changeEmail("bill@hotmail.com")
	bill.notify()

	// Pointers of type user can also be used to call methods
	// declared with both value and pointer receiver.
	joan := &user{"Joan", "joan@email.com"}
	joan.changeEmail("joan@hotmail.com")
	joan.notify() 													//(*joan).notify()

	// Create a slice of user values with two users.
	users := []user{
		{"ed", "ed@email.com"},
		{"erick", "erick@email.com"},
	}

	// Iterate over the slice of users switching
	// semantics. Not Good!
	for _, u := range users {
		u.changeEmail("")
	}

	fmt.Println(users)
}
