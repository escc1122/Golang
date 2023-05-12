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

	for typeOf.Kind() == reflect.Pointer {
		fmt.Println("pointer: ", valueOf.Pointer())
		typeOf = typeOf.Elem()
		valueOf = valueOf.Elem()
		fmt.Println("Elem() typeOf.Kind:", typeOf.Kind())
		fmt.Println("Elem() valueOf:", valueOf)
	}
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

// 測試 Slice Range 某些情況下與預期不一樣
// https://github.com/golang/gofrontend/blob/e387439bfd24d5e142874b8e68e7039f74c744d7/go/statements.cc#L5384
func testSliceRange() {
	type T struct {
		id int
	}
	t1 := T{id: 1}
	t2 := T{id: 2}
	ts1 := []T{t1, t2}
	ts2 := []*T{}
	ts3 := []*T{&t1, &t2}
	ts4 := []*T{}

	for _, t := range ts1 {
		//typeOf.Kind: ptr
		//valueOf: &{1}
		//pointer:  824633835688
		//Elem() typeOf.Kind: struct
		//Elem() valueOf: {1}
		//==============================================================
		//typeOf.Kind: ptr
		//valueOf: &{2}
		//pointer:  824633835688
		//Elem() typeOf.Kind: struct
		//Elem() valueOf: {2}
		reflectPrint(&t)
		ts2 = append(ts2, &t)
		fmt.Println("==============================================================")
	}
	//2
	//2
	for _, t := range ts2 {
		fmt.Println((*t).id)
	}

	fmt.Println("==============================================================")

	ts2 = []*T{}
	for _, t := range ts1 {
		tmp := t
		//typeOf.Kind: ptr
		//valueOf: &{1}
		//pointer:  824633835784
		//Elem() typeOf.Kind: struct
		//Elem() valueOf: {1}
		//==============================================================
		//typeOf.Kind: ptr
		//valueOf: &{2}
		//pointer:  824633835816
		//Elem() typeOf.Kind: struct
		//Elem() valueOf: {2}
		reflectPrint(&tmp)
		ts2 = append(ts2, &tmp)
		fmt.Println("==============================================================")
	}
	//1
	//2
	for _, t := range ts2 {
		fmt.Println((*t).id)
	}
	fmt.Println("==============================================================")

	for _, t := range ts3 {
		//typeOf.Kind: ptr
		//valueOf: &{1}
		//pointer:  824634400856
		//Elem() typeOf.Kind: struct
		//Elem() valueOf: {1}
		//==============================================================
		//typeOf.Kind: ptr
		//valueOf: &{2}
		//pointer:  824634400880
		//Elem() typeOf.Kind: struct
		//Elem() valueOf: {2}
		reflectPrint(t)
		ts4 = append(ts4, t)
		fmt.Println("==============================================================")
	}

	//1
	//2
	for _, t := range ts4 {
		fmt.Println((*t).id)
	}
	fmt.Println("==============================================================")

	// test
	t3 := T{id: 3}
	t3Point := &t3
	//typeOf.Kind: ptr
	//valueOf: &{3}
	//pointer:  824633835912
	//Elem() typeOf.Kind: struct
	//Elem() valueOf: {3}
	reflectPrint(t3Point)
	fmt.Println("==============================================================")
	t3 = T{id: 4}
	//typeOf.Kind: ptr
	//valueOf: &{4}
	//pointer:  824633835912
	//Elem() typeOf.Kind: struct
	//Elem() valueOf: {4}
	reflectPrint(t3Point)
	fmt.Println("==============================================================")
}

func main() {
	testSliceRange()
}
