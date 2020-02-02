package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// START OMIT
func main() {
	slice := make([]string, 0, 3)
	inspect(slice)
	slice = append(slice, "1")
	slice = append(slice, "2")
	slice = append(slice, "3")
	inspect(slice)
	slice = append(slice, "4")
	inspect(slice)
}
// END OMIT

func inspect(s []string) {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	fmt.Printf("slice address (%p), len(%d), cap(%d)\narray address (%d)\n\n", &s, len(s), cap(s), sh.Data)
}

