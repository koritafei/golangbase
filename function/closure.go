package main

import (
	"fmt"
)

func test(x int) func() {
	return func() {
		fmt.Println(x)
	}
}

func main() {
	t := test(123)
	t()
}
