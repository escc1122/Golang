package main

import (
	"fmt"
	"reflect"
)

type Address struct {
}

// getAddr 取得記憶體位置
func (a *Address) getAddr(x interface{}) int {
	v := reflect.ValueOf(x)
	return int(v.Pointer())
}
func (a *Address) getAddr2(x interface{}) string {
	addr := fmt.Sprintf("%p", x)
	return addr
}
