package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumIntsOrFloats(t *testing.T) {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	sum := SumIntsOrFloats(ints)
	assert.Equal(t, int64(46), sum)
}
