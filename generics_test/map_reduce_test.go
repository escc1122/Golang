package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	s := []int{1, 2, 3}

	filter := Filter(s, func(i int) bool { return i%2 == 0 })

	assert.Equal(t, 1, len(filter))
	assert.Equal(t, 2, filter[0])

}

func TestMap(t *testing.T) {
	s := []int{1, 2, 3}

	floats := Map(s, func(i int) float64 { return float64(i) })
	fmt.Println(floats)
}

func TestReduce(t *testing.T) {
	s := []int{1, 2, 3}

	sum := Reduce(s, 0, func(i, j int) int { return i + j })

	fmt.Println(sum)

}

func TestMapReduce(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7}

	ss := MapReduce[int]{s}

	sum := ss.Filter(
		func(i int) bool {
			return i%2 == 0
		}).
		Map(func(i int) int {
			return i * 5
		}).
		Reduce(func(i, j int) int {
			return i + j
		}, 0)

	assert.Equal(t, 60, sum)
}
