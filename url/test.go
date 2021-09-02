package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
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

	t, err := strconv.ParseInt("", 10, 64)

	fmt.Printf("t %v, err %v\n", t, err)

}
