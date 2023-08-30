package test1

import (
	"embed"
	"fmt"
)

//go:embed data/hello.txt
var f embed.FS

func embedTest() {
	data, err := f.ReadFile("data/hello.txt")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(string(data))
}
