package example1

import "fmt"

// START OMIT
type printer interface {
	print()
}

type user struct {
	name string
}

func (u user) print() {
	fmt.Println("User Name:", u.name)
}

func main() {
	u := user{"Bill"}
	entities := []printer{u, &u}

	u.name = "Bill_CHG"

	for _, e := range entities {
		e.print()
	}
}

// END OMIT
