package main

import "fmt"

func print(a interface{}) {
	switch a.(type) {
	case map[string][]string:
		for k, v := range a.(map[string][]string) {
			fmt.Printf("key: %v, value: %v\n", k, v)
		}
	case string:
		fmt.Printf("string %v\n", a.(string))
	case []string:
		for _, v := range a.([]string) {
			fmt.Printf(" value: %v\n", v)
		}
	}
}

func main() {
	var t = make(map[string][]string)
	t["qwe"] = []string{"sdfa", "adfsadf"}
	print(t)
	s := "string"
	print(s)

	ss := make([]string, 0)
	ss = append(ss, "1")
	ss = append(ss, "2")

	print(ss)
}
