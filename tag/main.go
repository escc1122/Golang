package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

var user = &User{"test1", "aaa@gmail.com"}

func main() {
	//test1()
	//test2()
	user.test3()
	test3()
}

func test1() {

	s := reflect.TypeOf(user).Elem()
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag)
	}
}

func test2() {
	data := []byte(`{"name" : test2 , "email" : "bbb@gmail.com"}`)
	json.Unmarshal(data, &user)
	jsondata, _ := json.Marshal(user)
	fmt.Println(string(jsondata))
}

type re struct {
	key string
	re  *[]re
}

func (u *User) test3() {
	//fmt.Printf("我是 %s, 谁又在调用我?\n", printMyName())
	printMyName()
}

func test3() {
	//fmt.Printf("我是 %s, 谁又在调用我?\n", printMyName())
	printMyName()
}

func printMyName() {
	pc, _, _, _ := runtime.Caller(1)
	//return runtime.FuncForPC(pc).Name()
	fmt.Println(runtime.FuncForPC(pc).Name())

}
