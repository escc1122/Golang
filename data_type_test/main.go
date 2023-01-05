package main

import (
	"fmt"
	"reflect"
)

func reflectPrint(test interface{}) {
	typeOf := reflect.TypeOf(test)
	valueOf := reflect.ValueOf(test)
	fmt.Println("typeOf.Kind:", typeOf.Kind())
	fmt.Println("valueOf:", valueOf)
}

func printPoint(msg string, test interface{}) {
	valueOf := reflect.ValueOf(test)
	fmt.Println(msg, "pointer:", valueOf.Pointer())
}

func testArray() {
	a := [2]string{"5", "6"}
	//[5 6]
	fmt.Println(a)

	func(funIn [2]string) {
		funIn[1] = "7"
	}(a)
	//[5 6]
	fmt.Println(a)

	func(funIn *[2]string) {
		funIn[1] = "7"
	}(&a)

	//[5 7]
	fmt.Println(a)

	string1 := "5"
	string2 := "6"

	//[0xc00003e250 0xc00003e260]
	b := [2]*string{&string1, &string2}
	fmt.Println(b)
	func(funIn [2]*string) {
		string3 := "7"
		funIn[1] = &string3
	}(b)

	//[0xc00003e250 0xc00003e260]
	fmt.Println(b)
}

func testSlice() {
	a := make([]string, 2)
	a[0] = "5"
	a[1] = "6"

	//typeOf.Kind: slice
	//valueOf: [5 6]
	reflectPrint(a)

	func(funIn []string) {
		funIn[1] = "7"
	}(a)
	//typeOf.Kind: slice
	//valueOf: [5 7]
	reflectPrint(a)
}

func testMap() {
	a := make(map[string]string, 5)
	a["5"] = "5"
	a["6"] = "6"
	//typeOf.Kind: map
	//valueOf: map[5:5 6:6]
	reflectPrint(a)
	func(funIn map[string]string) {
		funIn["7"] = "7"
	}(a)
	//typeOf.Kind: map
	//valueOf: map[5:5 6:6 7:7]
	reflectPrint(a)
}

func testStrut() {
	type a struct {
		id1 string
		id2 string
	}

	b := a{"5", "6"}
	//typeOf.Kind: struct
	//valueOf: {5 6}
	reflectPrint(b)
	func(funIn a) {
		funIn.id1 = "7"
	}(b)
	//typeOf.Kind: struct
	//valueOf: {5 6}
	reflectPrint(b)

	func(funIn *a) {
		funIn.id1 = "7"
	}(&b)
	//typeOf.Kind: ptr
	//valueOf: &{7 6}
	reflectPrint(&b)
}

func testFunc() {
	var a = func() {
		//do nothing
	}
	//typeOf.Kind: func
	//valueOf: 0x6be6a0
	reflectPrint(a)
	func(funIn func()) {
		//typeOf.Kind: func
		//valueOf: 0x6be6a0
		reflectPrint(funIn)
	}(a)
}

func testChannel() {
	a := make(chan int, 1)
	//typeOf.Kind: chan
	//valueOf: 0xc0000160e0
	reflectPrint(a)
	func(funIn chan int) {
		funIn <- 1
		//typeOf.Kind: chan
		//valueOf: 0xc0000160e0
		reflectPrint(a)
	}(a)

	fmt.Println(<-a)
}

// https://github.com/golang/gofrontend/blob/e387439bfd24d5e142874b8e68e7039f74c744d7/go/statements.cc#L5384
func testSliceRange() {
	type T struct {
		id int
	}
	t1 := T{id: 1}
	t2 := T{id: 2}
	ts1 := []T{t1, t2}
	ts2 := []*T{}
	for i, t := range ts1 {
		//test1 pointer: 824633835656
		//test1 pointer: 824633835656
		printPoint("test1", &t)

		//test2 pointer: 824633835680
		//test2 pointer: 824633835688
		printPoint("test2", &ts1[i])
		ts2 = append(ts2, &t)
	}

	//2
	//2
	for _, t := range ts2 {
		fmt.Println((*t).id)
	}
}

func main() {
	testSliceRange()
}
