package main

import (
	"fmt"
	"testing"
)

func TestWowStruct(t *testing.T) {
	b := &WowStruct[int, []int]{
		Data:     []int{5, 6, 7},
		MaxValue: 10,
		MinValue: 5,
	}
	fmt.Println(b)
}

func TestData_addData(t *testing.T) {
	data1 := []int32{10, 20, 30, 40, 50}
	data2 := []float32{10.1, 20.2, 30.3, 40.4, 50.5}
	d1 := Data[int32]{}
	d2 := Data[float32]{}
	d1.addData(data1...)
	d2.addData(data2...)
	sum1 := d1.sum()
	sum2 := d2.sum()
	fmt.Printf("sum1: %v (%T)\n", sum1, sum1)
	fmt.Printf("sum2: %v (%T)\n", sum2, sum2)
}
