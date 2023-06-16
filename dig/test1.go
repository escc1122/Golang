package main

import (
	"fmt"
)

type CarA struct {
}

func (r *CarA) Get() string {
	return "CarA"
}

func getCarA() *CarA {
	return &CarA{}
}

type CarB struct {
}

func (c *CarB) Get() string {
	return "CarB"
}

func getCarB() *CarB {
	return &CarB{}
}

func test1(carA *CarA, carB *CarB) {
	fmt.Println(carA.Get())
	fmt.Println(carB.Get())
}
