package main

import "fmt"

func pointerParam(name *int) {
	fmt.Printf("func address:%v value: %v\n", &name, name)
	*name = 200
	fmt.Printf("func address:%v value: %v\n", &name, name)
}

func main() {
	var a int
	a = 100
	b := &a
	fmt.Printf("main address:%v value: %v\n", &b, *b)
	pointerParam(b)
	fmt.Printf("main address:%v value: %v\n", &b, *b)
}
