package main

import "fmt"

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method notifies users of different events.
func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

type admin struct {
	user  // Embedded Type
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "Rob Pike",
			email: "robpike@google.com",
		},
		level: "superuser",
	}

	// We can access the inner type's method directly.
	ad.user.notify()

	// Because of inner type promotion, we can access the notify method directly through the outer
	// type.
	ad.notify()

	// We are passing the address of outer type value. Because of inner type promotion, the outer
	// type now implements all the same contract as the inner type.
	sendNotification(&ad)
}

type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}
