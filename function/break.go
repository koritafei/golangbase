package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%5 == 0 {
			fmt.Printf("%d\n", i)
			break
		}
		fmt.Printf("i = %d\n", i)
	}
	fmt.Printf("End Main")
}
