package main

import (
	"fmt"
	"reflect"
)

type IPerson interface {
	test()
}

type Person struct {
	Name string `bson:"name"`
	Age  int    `bson:"age"`
}

func (p *Person) test() {
	fmt.Printf("person Name: %s, Age: %d \n", p.Name, p.Age)
}

func getStrutValue() {
	//p := Person{Name: "John", Age: 30}
	p := &Person{Name: "John", Age: 30}
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	for v.Kind() == reflect.Pointer {
		t = t.Elem()
		v = v.Elem()
		fmt.Println("Value is Pointer")
	}

	fmt.Println("Type:", t.Name())
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		fmt.Printf("%s: %v\n", field.Name, value)
	}
}

// createStruct  使用 reflect.New 取得 person struct
func createStruct() {
	// 取得Person型別
	personType := reflect.TypeOf(Person{})
	// 實體化一個Person struct
	person := reflect.New(personType).Elem().Interface().(Person)
	// 初始化Person struct的屬性值
	person.Name = "John"
	person.Age = 30
	// 輸出Person struct的屬性值
	fmt.Printf("person Name: %s, Age: %d \n", person.Name, person.Age)
}

// createStruct2  使用 reflect.New 取得 person struct
func createStruct2() {
	// 取得Person型別
	personType := reflect.TypeOf(Person{})
	personPtr := reflect.New(personType)
	fmt.Printf("person reflect.New kind: %s \n", personPtr.Kind())
	personStruct := personPtr.Elem()
	fmt.Printf("person New Elem kind: %s \n", personStruct.Kind())
	personStruct.FieldByName("Name").SetString("Tom")
	personStruct.FieldByName("Age").SetInt(20)
	person := personStruct.Interface().(Person)
	fmt.Printf("person2 Name: %s, Age: %d \n", person.Name, person.Age)
}

// createStruct3  使用 reflect.Zero 取得 person struct
func createStruct3() {
	// 取得Person型別
	personType := reflect.TypeOf(Person{})
	personValue := reflect.Zero(personType)
	fmt.Printf("person reflect.Zero kind: %s \n", personValue.Kind())
	person := personValue.Interface().(Person)
	fmt.Printf("person Name: %s, Age: %d \n", person.Name, person.Age)
}

// createStruct4  取得 person interface
func createStruct4() {
	// 取得Person型別
	personType := reflect.TypeOf(Person{})
	personPtr := reflect.New(personType)
	fmt.Printf("person2 reflect.New kind: %s \n", personPtr.Kind())

	personStruct := personPtr.Elem()
	fmt.Printf("person2 New Elem kind: %s \n", personStruct.Kind())
	personStruct.FieldByName("Name").SetString("Tom")
	personStruct.FieldByName("Age").SetInt(20)

	person := personPtr.Interface().(IPerson)
	person.test()
}

func getTag() {
	reflectTest := Person{}
	t := reflect.TypeOf(reflectTest)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Println(field.Tag.Get("bson"))
	}
}
