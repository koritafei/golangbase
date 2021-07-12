package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	c := make(chan int)
	var send chan<- int = c
	var recv <-chan int = c

	go func() {
		for x := range recv {
			fmt.Printf("%+v\n", x)
		}

		wg.Done()
	}()

	go func() {
		for i := 0; i < 4; i++ {
			send <- i
		}

		wg.Done()
		close(send)
	}()

	wg.Done()
}
