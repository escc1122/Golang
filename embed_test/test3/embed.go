package test1

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed data/*
var f embed.FS

func embedTest() {
	data, err := f.ReadFile("data/hello.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))

	data2, err := f.ReadFile("data/hello2.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data2))
}
