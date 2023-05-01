package main

import (
	"fmt"
	"reflect"
)

func main() {
	// variable declaration with initial value
	var r rune = 'a'
	fmt.Printf("Value: %c, Unicode: %U, Type: %T\n", r, r, r)

	// Shorthand declaration and assignment
	a1 := 97
	a2 := 'b'
	a3 := '♬'

	// Printing the value, unicode equivalent and type of the variable
	fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", a1, a1, reflect.TypeOf(a1))
	fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", a2, a2, reflect.TypeOf(a2))
	fmt.Printf("Value: %c, Unicode: %U, Type: %s\n", a3, a3, reflect.TypeOf(a3))
}
