package main

import (
	"errors"
	"fmt"
)

// START OMIT
var (
	ErrBadRequest = errors.New("bad request")
)

func main() {
	if err := webCall(); err != nil {
		switch err {
		case ErrBadRequest:
			fmt.Println("bad request occurred")
			return
		}
	}
}

func webCall() error {
	return ErrBadRequest
}

// END OMIT
