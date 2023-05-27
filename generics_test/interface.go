package main

import (
	"strconv"
)

type MyGenericInterface[T Number] interface {
	DoSomething(T) string
}

type MyStruct struct {
}

func (s *MyStruct) DoSomething(i int) string {
	return "Received integer: " + strconv.Itoa(i)
}

// =================================================
