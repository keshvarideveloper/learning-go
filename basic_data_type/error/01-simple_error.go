package main

import (
	"fmt"
	"strconv"
)

func main() {
	_, err := strconv.Atoi("ABC123")
	if err != nil {
		// Print err message without custom message
		fmt.Println(err)
	}
}
