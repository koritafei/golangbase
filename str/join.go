package main

import (
	"bytes"
	"strings"
	"fmt"
)

func testJoin() string {
	s := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		s[i] = "a"
	}

	return strings.Join(s, "")
}

func testByteBuffer() string {
	var b bytes.Buffer
	b.Grow(1000)
	for i := 0; i < 1000; i++ {
		b.WriteString("a")
	}

	return b.String()
}

func main() {
	fmt.Println(testJoin())
	fmt.Println(testByteBuffer())
}
