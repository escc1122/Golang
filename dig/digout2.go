package main

import (
	"fmt"
	"go.uber.org/dig"
)

type digOut struct {
	dig.Out
	CarA *CarA
	CarB *CarB
}

func getCar4() digOut {
	return digOut{
		CarA: &CarA{},
		CarB: &CarB{},
	}
}

func test4(carA *CarA, carB *CarB) {
	fmt.Println(carA.Get())
	fmt.Println(carB.Get())
}

//func test4(out digOut) {
//	fmt.Println(out.CarA.Get())
//	fmt.Println(out.CarB.Get())
//}
