package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "ZnJvbToxMEIzMDk1MDEwO3dtOjkwMDZfMjAwMTtsdWljb2RlOjEwMDAwMDAzO3VpY29kZToxMDAwMDAwMztmaWQ6MTAwMTAzdHlwZT0xJnE96Ze66Jyc5aS05YOPJnQ9MTtsZmlkOjEwMDMwM3R5"
	fmt.Println(str)
	str1, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("%v", err)
	}
	fmt.Printf("%s \n", str1)

}
