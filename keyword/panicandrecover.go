package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("in main")

	go func() {
		defer fmt.Println("in goroutine")
		panic("")
	}()

	time.Sleep(time.Second * time.Duration(1))
}
