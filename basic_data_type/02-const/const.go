package main

import "fmt"

const (
	year     = 365
	leapYear = int32(366)
)

func main() {
	hours := 24
	minutes := int32(60)

	fmt.Println(hours * year)
	fmt.Println(minutes * year)
	fmt.Println(minutes * leapYear)

	// Can't assign new value to the const variable
	/*
		year = year + 1
		fmt.Println("year with new value: ", year) // Result: cannot assign to year
	*/
}
