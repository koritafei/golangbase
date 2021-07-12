package main

import (
	"fmt"
	"reflect"
)

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func main() {
	var m manager

	t := reflect.TypeOf(&m)

	fmt.Printf("Kind %+v\n", t.Kind())

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset)

		if f.Anonymous {
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Printf("%+v %+v\n", af.Name, af.Type)
			}
		}
	}

}
