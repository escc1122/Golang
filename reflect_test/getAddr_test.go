package main

import (
	"fmt"
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
