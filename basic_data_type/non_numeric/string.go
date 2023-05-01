package main

import "fmt"

func main() {

	var s string = "Hello界"

	// for loop
	for i := 0; i < len(s); i++ {
		fmt.Printf("Value: %c, Unicode: %U, Type: %T\n", s[i], s[i], s[i])
	}

	// range
	for index, runeValue := range s {
		fmt.Printf("%#U starts at byte position %d type %T \n", runeValue, index, runeValue)
	}
}
