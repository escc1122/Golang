package main

import (
	"fmt"

	"go.uber.org/dig"
)

type digIn struct {
	dig.In
	CarA *CarA
	CarB *CarB
}

func test2(in digIn) {
	fmt.Println(in.CarA.Get())
	fmt.Println(in.CarB.Get())
}
