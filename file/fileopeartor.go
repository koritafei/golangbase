/**
 * @file: fileopeartor.go
 * @description:
 * @author: koritafei
 * @create: 2021-05-08 17:30
 * @version v0.1
 * */

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(10)
	ch := make(chan string, 1000)

	filename := "/Users/koritafei/Desktop/code/golangbase/file/fileopeartor.go"
	readFile(filename, ch)

	close(ch)


	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("continue %d\n", i)
			processChannel(&wg, ch)
		}(i)
	}

	wg.Wait()
}

func readFile(filename string, ch chan string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	br := bufio.NewReader(f)

	for {
		line, _, c := br.ReadLine()
		if c == io.EOF {
			return
		}
		ch <- string(line)
	}

}

func processChannel(wg *sync.WaitGroup, ch chan string) {
	fmt.Printf("channel len %d\n", len(ch))
	for val := range ch {
		fmt.Println(val)
	}
	wg.Done()
}
