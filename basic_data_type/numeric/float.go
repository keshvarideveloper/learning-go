package main

import "fmt"

var (
	// All of signed integer can be bigger and smaller than 0
	float32Var float32 = -133.2565
	float64Var float64 = 5248.55

	// Type can be omitted and its type will be float64
	floatVar = 25.025
)

func main() {

	// Print intVar type and value
	fmt.Printf("type of float32Var is: %T, and the value is: %f\n", float32Var, float32Var) // type of float32Var is: float32, and the value is: -133.256500

	// Print intVar type and value
	fmt.Printf("type of float64Var is: %T, and the value is: %f\n", float64Var, float64Var) // type of float64Var is: float64, and the value is: 5248.550000

	// Print intVar type and value
	fmt.Printf("type of floatVar is: %T, and the value is: %f", floatVar, floatVar) // type of floatVar is: float64, and the value is: 25.025000

}
