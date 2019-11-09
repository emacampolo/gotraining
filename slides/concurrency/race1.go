package main

import "fmt"

//START OMIT
func main() {
	var data int

	go func() {
		data++
	}()

	if data == 0 {
		fmt.Printf("the value is %v", data)
	}
}

//END OMIT
