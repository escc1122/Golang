package main

import (
	"fmt"
)

type CarA struct {
}

func (r *CarA) Get() string {
	return "CarA"
}

type CarB struct {
}

func (c *CarB) Get() string {
	return "CarB"
}

func getCarA() *CarA {
	return &CarA{}
}

func getCarB() *CarB {
	return &CarB{}
}

func test1(carA *CarA, carB *CarB) {
	fmt.Println(carA.Get())
	fmt.Println(carB.Get())
}
