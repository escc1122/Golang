package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddress_getAddr(t *testing.T) {
	address := Address{}

	address1 := &address

	address2 := &address

	address3 := &*address2

	fmt.Println(address.getAddr(address1))
	fmt.Println(address.getAddr(address2))
	fmt.Println(address.getAddr(address3))

	fmt.Println(address.getAddr2(address1))
	fmt.Println(address.getAddr2(address2))
	fmt.Println(address.getAddr2(address3))
}

func TestAddress_getAddr2(t *testing.T) {
	address := &Address{}
	testMap := make(map[string]string)
	fmt.Println("reflect.TypeOf(testMap).Kind(): ", reflect.TypeOf(testMap).Kind())
	testMap2 := &testMap
	fmt.Println("reflect.TypeOf(testMap).Kind(): ", reflect.TypeOf(testMap2).Kind())

	func(inTestMap map[string]string) {
		fmt.Println("in: ", address.getAddr(testMap))
	}(testMap)

	fmt.Println("out: ", address.getAddr(testMap))
}

// 淺層複製
func TestAddress_getAddr3(t *testing.T) {
	//address := Address{}
	p1 := Person{Name: "John", Age: 30}
	p2 := p1 // 複製 p1 到 p2

	p2.Name = "test"

	fmt.Println(&p1) // &{John 30}
	fmt.Println(&p2) // &{test 30}

	p3 := &Person{Name: "John", Age: 30}
	p4 := p3

	p4.Name = "test5"

	fmt.Println(p3) // &{test5 30}
	fmt.Println(p4) // &{test5 30}
}
