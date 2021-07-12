/*
 * @Author: your name
 * @Date: 2021-03-05 12:09:58
 * @LastEditTime: 2021-04-12 13:33:01
 * @LastEditors: your name
 * @Description: In User Settings Edit
 * @FilePath: /golangbase/str/stringop.go
 */
package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func ToString(s []byte) string {
	return *(*string)(unsafe.Pointer(&s))
}

func main() {
	bs := []byte("hello world")
	s := ToString(bs)
	fmt.Println(s)

	str := "__TS__{TS}"
	r := strings.NewReplacer("__TS__", "100", "{TS}", "100")
	fmt.Printf(r.Replace(str))

}
