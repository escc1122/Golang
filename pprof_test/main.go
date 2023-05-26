package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	_, err := fmt.Println("Hello, ithome")
	if err == nil {
		gorace()
	}
}

func gorace() {
	c := make(chan bool)
	m := make(map[string]string)
	go func() {
		m["1"] = "a" // First conflicting access.
		c <- true
	}()
	m["2"] = "b" // Second conflicting access.
	<-c
	for k, v := range m {
		fmt.Println(k, v)
	}

	http.ListenAndServe("localhost:6060", nil)
}
