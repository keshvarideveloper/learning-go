package main

import (
	"fmt"
)

func main() {
	a, b := 0, 0

	err := customErrMsgWithFmtError(a, b)
	if err != nil {
		fmt.Println(err)
	}
}

// Custom error message with errors.New()
func customErrMsgWithFmtError(a int, b int) error {
	if a == 0 && b == 0 {

		return fmt.Errorf("a %d and b %d", a, b)
	}

	return nil
}
