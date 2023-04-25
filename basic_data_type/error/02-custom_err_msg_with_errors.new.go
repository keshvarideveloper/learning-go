package main

import (
	"errors"
	"fmt"
)

func main() {
	a, b := 0, 0

	err := customErrMsgWithErrorsNew(a, b)
	if err != nil {
		fmt.Println(err)
	}
}

// Custom error message with errors.New()
func customErrMsgWithErrorsNew(a int, b int) error {
	if a == 0 && b == 0 {

		return errors.New("this is a custom error message")
	}

	return nil
}
