package main

import "fmt"

func test(f func()) {
	f()
}

func test2() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func main() {
	// 直接调用
	func(s string) {
		fmt.Println(s)
	}("Hello world")

	// 赋值到变量

	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 2))

	// 函数参数
	test(func() {
		fmt.Println("hello function paramer")
	})

	// 函数返回值
	a := test2()
	fmt.Println(a(2, 3))
}
