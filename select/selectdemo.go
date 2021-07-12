package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	t := time.NewTicker(time.Duration(1) * time.Second)

	go func() {
		i := 0
		for {
			ch <- i
			i++
			time.Sleep(time.Second * time.Duration(1))
		}
	}()

	for {
		select {
		case i := <-ch:
			fmt.Println(i)
		case <-t.C:
			fmt.Println("Time Ticker")
		default:
			time.Sleep(time.Second * time.Duration(1))
			fmt.Println("Default")
		}
	}

}
