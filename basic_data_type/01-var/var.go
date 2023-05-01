package main

import "fmt"

// Declaring multiple variables at once
var (
	x    int                  // assign zero value
	y           = 20          // set type according what is the value. for example in this sample type is int
	z    int    = 30          // explicit type, and an assignment
	d, e        = 40, "hello" // multiple variables as different types
	f, g string               // multiple variables as same types

)

func main() {
	fmt.Printf("x: %v, y: %v, z: %v, d: %v, e: %v, f: %v, g: %v\n",
		x, y, z, d, e, f, g)

	// 	Short declaration with :=
	t := true
	fmt.Println("t: ", t)

	// := allows you to assign values to existing variables, too. As long as there is one new variable on the lefthand side of the :=
	t, s := false, "Hello"
	fmt.Printf("new value of t is: %v , value of new variable 's' is: %v ", t, s)

}
