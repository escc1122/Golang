package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
}

func main() {

	a := &Person{
		firstName: "Alex",
		lastName:  "Anderson",
	}

	b := &Person{
		firstName: "al",
		lastName:  "test",
	}

	c := Person{
		firstName: "al",
		lastName:  "test2",
	}

	d := Person{
		firstName: "al",
		lastName:  "test3",
	}

	fmt.Println(a)
	fmt.Println(a.lastName)
	fmt.Println(b)
	fmt.Println(b.lastName)
	fmt.Println(c)
	fmt.Println(c.lastName)
	fmt.Println(d)
	fmt.Println(d.lastName)
}
