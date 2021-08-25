package main

import (
	"fmt"

	version "github.com/hashicorp/go-version"
)

func main() {
	strios := "os14.5.0"
	strandroid := "android8.0.0"

	fmt.Println("os14.5.1" >= strios)
	fmt.Println("android8.0.0" >= strandroid)
	fmt.Println("os24.5.1" >= strios && len("os24.5.1") >= len(strios))

	fmt.Println("android12.0.0" >= strandroid)

	fmt.Println("android12.0.0" >= strandroid && len("android12.0.0") >= len(strandroid))

	strandroidt := strandroid[7:len(strandroid)]

	v1, _ := version.NewVersion(strandroidt)
	v2, _ := version.NewVersion("12.0.0")
	if v1.LessThan(v2) {
		fmt.Printf("%s is less than %s\n", v1, v2)
	}

}
