package main

import "fmt"

func main() {
	var b []int
	for i := 1; i < 9; i++ {
		b = append(b, i)
	}
	a := b[0:5]
	fmt.Println(a)

	// integer slice
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	//array
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // [2 3 5 7 11 13]
}
