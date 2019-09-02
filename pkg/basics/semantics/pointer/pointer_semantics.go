// Sample program to show the basic concept of using a pointer
// to share data.
package main

import (
	`fmt`
)

func main() {

	// Declare variable of type int with a value of 10.
	count := 10

	// Display the "value of" and "address of" count.
	fmt.Println("count: Value of", count, "Addr of", &count)

	// Pass the "address of" count. We still pass by value but it happens to be an address.
	// This is not pass by reference!
	increment(&count)

	fmt.Println("count: Value of", count, "Addr of", &count)
}

// increment declares inc as a pointer variable whose value is
// always an address and points to values of type int.
func increment(inc *int) {

	// The (*) operator declares a pointer variable and the "Value that the pointer points to"
	// Indirect read, modify, write operation.
	// We are having side effects since we no longer acts in a sandbox.
	*inc++
	fmt.Println("inc: Value of", inc, "Addr of", &inc, "Value Points To", *inc)
}

/*
#include <stdio.h>

void swap(int &i, int &j) {
  int temp = i;
  i = j;
  j = temp;
}

int main(void) {
  int a = 10;
  int b = 20;

  swap(a, b);
  printf("A is %d and B is %d\n", a, b); // A is 20 and B is 10
  return 0;
}
 */