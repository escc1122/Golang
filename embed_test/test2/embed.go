package test1

import (
	_ "embed"
	"fmt"
)

//go:embed data/hello.txt
var s string

func embedTest() {
	fmt.Print(s)
}
