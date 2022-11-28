package main

import (
	"fmt"
	"reflect"
)

type IReflectTest interface {
	test()
}

type reflectTest struct {
	Aaaa string `bson:"aaaa"` //要大寫 不然反射不到
	Bbbb string `bson:"bbbb"`
}

func (r *reflectTest) test() {
	fmt.Println("IReflectTest")
}

type reflectTest2 struct {
	Aaaa string `bson:"aaaa"`
	Bbbb string `bson:"bbbb"`
}

func (r reflectTest2) test() {
	fmt.Println("IReflectTest")
}

func testSetValue() {
	reflectTest := reflectTest{}
	t := reflect.TypeOf(reflectTest)
	v := reflect.ValueOf(&reflectTest).Elem()
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i).Name)
		key := t.Field(i).Name
		f := v.FieldByName(key)
		f.SetString("55555" + key)
	}
	fmt.Println(reflectTest.Aaaa)
	fmt.Println(reflectTest.Bbbb)
}

func testTag() {
	reflectTest := reflectTest{}
	t := reflect.TypeOf(reflectTest)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Tag.Get("bson"))
	}
}

func test111(test interface{}) {
	typeOf := reflect.TypeOf(test)
	valueOf := reflect.ValueOf(test)

	fmt.Println("type:", typeOf)
	fmt.Println("typeOf.Kind:", typeOf.Kind())
	if typeOf.Kind() == reflect.Pointer {
		typeOf = typeOf.Elem()
		fmt.Println("typeOf.Elem():", typeOf)
		fmt.Println("typeOf.Elem().Kind:", typeOf.Kind())
	}

	fmt.Println("value:", valueOf)
	fmt.Println("valueOf.Kind:", valueOf.Kind())
	if valueOf.Kind() == reflect.Pointer {
		valueOf = valueOf.Elem()
		fmt.Println("valueOf.Elem():", valueOf)
		fmt.Println("valueOf.Elem().Kind:", valueOf.Kind())
	}

	fmt.Println("=========================================")
}

func test() {
	test111(reflectTest{})
	test111(&reflectTest{})
	var test IReflectTest = &reflectTest{}
	var test2 IReflectTest = reflectTest2{}
	test111(test)
	test111(test2)

}

func main() {
	testSetValue()
}
