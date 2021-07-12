package main

import (
	"fmt"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(1)
	exit := make(chan struct{})

	go func() {
		go func() {
			fmt.Printf("b\n")
		}()

		for i := 0; i < 4; i++ {
			fmt.Println("a: ", i)
			if i == 1 {
				runtime.Gosched() // 让出协程执行权限
			}
		}

		close(exit)
	}()

	<-exit

}
