package main

import (
	"strconv"
	"strings"
)

type IString interface {
	ToString() string
}

type String1 struct {
}

func (s *String1) ToString() string {
	return "string1"
}

type String2 struct {
}

func (s *String2) ToString() string {
	return "string2"
}

type myint int

func (i myint) ToString() string {
	return strconv.Itoa(int(i))
}

func Stringify[T IString](s []T) (ret []string) {
	for _, v := range s {
		ret = append(ret, v.ToString())
	}
	return ret
}

func Stringify2[T IString](s []T) string {

	var ret []string

	for _, v := range s {
		ret = append(ret, v.ToString())
	}

	return strings.Join(ret, ",")
}
