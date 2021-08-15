package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	str := "aaa^ssss^ccc^aaaaaa"
	fmt.Println(str)
	v, e := json.Marshal(str)
	if e != nil {
		panic(e)
	}

	fmt.Printf("%v\n", v)
	fmt.Printf("%v\n", string(v))

	a := ""
	e = json.Unmarshal(v, &a)

	if e != nil {
		panic(e)
	}

	fmt.Printf("%v\n", a)

	arr := strings.Split(string(a), "^")

	fmt.Printf("arr %v\n", arr)

	return
}
