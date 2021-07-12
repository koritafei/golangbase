package main

import "fmt"

func testStruct() {
	type calc struct {
		mul func(x, y int) int
	}

	x := calc{
		mul: func(x, y int) int {
			return x - y
		},
	}

	fmt.Println(x.mul(1, 2))
}

func testChannels() {
	c := make(chan func(int, int) int, 2)
	c <- func(x, y int) int { return x + y }

	fmt.Println((<-c)(1, 2))
}

func main() {
	testStruct()
	testChannels()

}
