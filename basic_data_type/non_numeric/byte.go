package main

import "fmt"

func main() {
	// assign by decimal integer (base 10)
	var ch byte = 90
	fmt.Printf("%08b %02x %v\n", ch, ch, ch)

	// assign an ASCII character
	var ch0 byte = 'Z'
	fmt.Printf("%08b %02x %v\n", ch0, ch0, ch0)

	// assign by different bases:
	var ch1 byte = 0b01011010 // Binary
	var ch2 byte = 0o132      // Octal
	var ch3 byte = 0x5a       // heX

	fmt.Printf("%08b %02x %v\n", ch1, ch1, ch1)
	fmt.Printf("%08b %02x %v\n", ch2, ch2, ch2)
	fmt.Printf("%08b %02x %v\n", ch3, ch3, ch3)
}
