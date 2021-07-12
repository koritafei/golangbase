/**
 * @file: msgpackDemo.go
 * @description:
 * @author: koritafei
 * @create: 2021-05-11 17:01
 * @version v0.1
 * */

package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
)

type ts struct {
	Bhv      string `msgpack:"Bhv"`
	Mac      string `msgpack:"X"`      // mac
	Brand    string `msgpack:"B"`      // brand
	Idfa     string `msgpack:"D"`      // idfa
	Os       string `msgpack:"O"`      // os版本
	Oaid     string `msgpack:"A"`      // oaid
	UA       string `msgpack:"UA"`     // user agent
	Android  string `msgpack:"AndId"`  // android id
	CaId     string `msgpack:"CaId"`   // ios 14+
	GxCaId   string `msgpack:"GxCaId"` // 中广协caid
	Model    string `msgpack:"M"`      // model
	Language string `msgpack:"L"`      // language
	AAid     string `msgpack:"AA"`     // aaid
}

func main() {
	t := &ts{
		Bhv:   "Bhv",
		Mac:   "kkk",
		Brand:   "ttt",
		Idfa: "maxmax",
		Os:"ios",
		Oaid: "8267ACAF-FEC4-AA2A",
		UA:"8267ACAF-FEC4-AA2A",
		Android: "8267ACAF-FEC4-AA2A",
		CaId: "8267ACAF-FEC4-AA2A",
		GxCaId: "8267ACAF-FEC4-AA2A",
		Model: "xiaomi",
		Language: "zh-cn",
		AAid:  "8267ACAF-FEC4-AA2A-4678-330B6B935478",
	}

	b, err := msgpack.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", b)

	var out = ts{}

	err = msgpack.Unmarshal(b, &out)
	if err != nil {
		panic(err)
	}
	fmt.Println(out)

}
