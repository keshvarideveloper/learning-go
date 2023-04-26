package main

import "fmt"

var (
	// All of signed integer can be bigger and smaller than 0
	int8var  int8  = -128
	int16Var int16 = 32767
	int32Var int32 = -2147483648
	int64var int64 = 9223372036854775807

	// Can be either 32 or 64 bits each—their size is defined by Go itself
	intVar = 10

	// All of unsigned integer can be only bigger or equal  0
	uint8var  uint8  = 255
	uint16Var uint16 = 65535
	uint32Var uint32 = 4294967295
	uint64var uint64 = 18446744073709551615

	// Can be either 32 or 64 bits each—their size is defined by Go itself
	uintVar uint = 10
)

func main() {
	fmt.Println("minimum number of int8 is:", int8var)
	fmt.Println("maximum number of int16 is:", int16Var)
	fmt.Println("minimum number of int32 is:", int32Var)
	fmt.Println("maximum number of int64 is:", int64var)

	// Print intVar type and value
	fmt.Printf("type of intVar is: %T, and the value is: %d", intVar, intVar) // type of intVar is: int, and the value is: 10

	fmt.Println("minimum number of uint8 is:", uint8var)
	fmt.Println("maximum number of uint16 is:", uint16Var)
	fmt.Println("minimum number of uint32 is:", uint32Var)
	fmt.Println("maximum number of uint64 is:", uint64var)

	// Print intVar type and value
	fmt.Printf("type of uintVar is: %T, and the value is: %d", uintVar, uintVar) // type of intVar is: uint, and the value is: 10

}
