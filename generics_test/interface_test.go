package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyStruct_DoSomething(t *testing.T) {
	var obj MyGenericInterface[int]
	obj = &MyStruct{}
	assert.Equal(t, "Received integer: 42", obj.DoSomething(42))
}
