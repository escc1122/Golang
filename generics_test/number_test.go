package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sum(t *testing.T) {
	n1, n2 := 5, 6

	sumValue := sum(n1, n2)
	if sumValue != 11 {
		t.Errorf("sum() = %v, want %v", sumValue, 11)
	}

	n3, n4 := int64(5), int64(11)

	sumValue2 := sum(n3, n4)
	if sumValue2 != 16 {
		t.Errorf("sum() = %v, want %v", sumValue2, 16)
	}
}

func Test_sum2(t *testing.T) {
	v := sum2(5, 7)
	assert.Equal(t, 12, v)
	v2 := sum2(int64(5), int64(7))
	assert.Equal(t, int64(12), v2)
}

func Test_sum3(t *testing.T) {
	var n1, n2 intA
	n1 = 5
	n2 = 7

	v := sum3(n1, n2)
	assert.Equal(t, intA(12), v)
}
