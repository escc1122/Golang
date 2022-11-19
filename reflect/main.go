package main

import (
	"fmt"
	"reflect"
)

type reflectTest struct {
	Aaaa string `bson:"aaaa"` //要大寫 不然反射不到
	Bbbb string `bson:"bbbb"`
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

func main() {
	testSetValue()
}
