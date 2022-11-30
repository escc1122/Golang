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
	//fmt.Println("IReflectTest")
	reflectPrint(r)
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

func reflectPrint(test interface{}) {
	typeOf := reflect.TypeOf(test)
	valueOf := reflect.ValueOf(test)

	fmt.Println("typeOf.Kind:", typeOf.Kind())
	fmt.Println("typeOf:", typeOf)
	i := 0
	for typeOf.Kind() == reflect.Pointer {
		typeOf = typeOf.Elem()
		fmt.Println(i, " typeOf.Elem().Kind:", typeOf.Kind())
		fmt.Println(i, " typeOf.Elem():", typeOf)
		i++
	}

	fmt.Println("valueOf.Kind:", valueOf.Kind())
	fmt.Println("valueOf:", valueOf)

	j := 0
	for valueOf.Kind() == reflect.Pointer {
		valueOf = valueOf.Elem()
		fmt.Println(j, " valueOf.Elem().Kind:", valueOf.Kind())
		fmt.Println(j, " valueOf.Elem():", valueOf)
	}

}

func (r *reflectTest) test1() {
	reflectPrint(r)
}

func test() {

	fmt.Println("========== print reflectTest{}")
	reflectPrint(reflectTest{})
	fmt.Println("========== print &reflectTest{}")
	reflectPrint(&reflectTest{})
	var test IReflectTest = &reflectTest{}
	var test2 IReflectTest = reflectTest2{}
	fmt.Println("========== print interface &reflectTest{}")
	reflectPrint(test)
	fmt.Println("========== print interface reflectTest{}")
	reflectPrint(test2)

	a := &reflectTest{}
	fmt.Println("========== print &reflectTest{} function")
	a.test1()

	fmt.Println("========== print c := &a ;a := &reflectTest{}")
	c := &a
	reflectPrint(c)

	fmt.Println("========== print reflectTest{} function")
	b := reflectTest2{}
	b.test()

}

func main() {
	test()
}
