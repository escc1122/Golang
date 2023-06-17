package main

import "fmt"

func getCar() (*CarA, *CarB) {
	return &CarA{}, &CarB{}
}

func test3(carA *CarA, carB *CarB) {
	fmt.Println(carA.Get())
	fmt.Println(carB.Get())
}
