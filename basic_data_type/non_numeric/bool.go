package main

import "fmt"

func main() {
	// variable declaration with initial value
	var a bool = false

	// Shorthand declaration and assignment
	b := true
	fmt.Printf("value of a is %v\nvalue of b is %v\n", a, b)

	c := a == b
	fmt.Printf("value of c is %v\n", c)

	// Logical Operators
	fmt.Println()
	fmt.Println("Logical Operators:")
	fmt.Println("a && b:", a && b)
	fmt.Println("a || b:", a || b)
	fmt.Println("!a:", !a)
	fmt.Println("!b:", !b)

	// Comparison Operators
	x := 5
	y := 8
	fmt.Println()
	fmt.Println("Comparison Operators:")
	fmt.Println("x == y:", x == y)
	fmt.Println("x != y:", x != y)
	fmt.Println("x < y:", x < y)
	fmt.Println("x > y:", x > y)
	fmt.Println("x <= y:", x <= y)
	fmt.Println("x >= y:", x >= y)

}
