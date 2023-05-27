package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ToString(t *testing.T) {
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

func TestStringify(t *testing.T) {
	x := []myint{myint(1), myint(2), myint(3)}
	test1 := Stringify(x)

	assert.Equal(t, "1", test1[0])
	assert.Equal(t, "2", test1[1])
	assert.Equal(t, "3", test1[2])

	y := []IString{&String1{}, &String2{}}

	test2 := Stringify2(y)
	assert.Equal(t, "string1,string2", test2)

}

func TestStringify2(t *testing.T) {

}
