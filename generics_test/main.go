package main

import (
	"fmt"
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

func main() {
	fmt.Println(sum(5, 7))
	fmt.Println(sum2(5, 7))

	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}
	fmt.Println(SumIntsOrFloats(ints))

	x := []myint{myint(1), myint(2), myint(3)}
	fmt.Println(Stringify(x))

	y := []IString{&String1{}, &String2{}}
	fmt.Println(Stringify2(y))

}

func sum[num int | int64](n1 num, n2 num) num {
	return n1 + n2
}

type Number interface {
	int | int64
}

func sum2[num Number](n1 num, n2 num) num {
	return n1 + n2
}

func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
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
