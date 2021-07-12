package main

import (
	"fmt"
	"reflect"
)

type S struct{}

type T struct {
	S
}

func (S) sVal() {
	fmt.Println("s val")
}

func (*S) sPtr() {
	fmt.Println("s pointer")
}

func (T) tVal() {
	fmt.Println("t val")
}

func (*T) tPtr() {
	fmt.Println("t pointer")
}

func methodSet(a interface{}) {
	t := reflect.TypeOf(a)
	n := t.NumMethod()
	for i := 0; i < n; i++ {
		m := t.Method(i)
		fmt.Println(m)
	}
}

func main() {
	var tt T

	methodSet(tt)
	fmt.Println("================================================================")
	methodSet(&tt)
}
