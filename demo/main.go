package main

import (
	"inter/data"
	"inter/inter"
)

func main() {
	s := data.Person{
		Name: "8267ACAF",
		Age:  12,
		Job:  "student",
	}

	i := inter.Student{}
	i.Printf(s)
}
