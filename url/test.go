package main

import (
	"fmt"
	"net/url"
	"regexp"
)

func main() {
	str := "中国 你好   +++---&&&0987654321！@#￥%……&*（）——+}{：『？》》《，。、__AD__ __asas__"

	fmt.Println(str)
	a := url.QueryEscape(str)
	fmt.Println(a)
	re, _ := regexp.Compile("\\_\\_\\w+\\_\\_")

	fmt.Println(re.MatchString(str))
	str = re.ReplaceAllString(str, "sdfsdf")
	fmt.Println(str)
}
