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

// 研究中 不跑
func TestMyStruct_ITest(t *testing.T) {
	nodes := make([]TreeEntity[string, *TestAAA], 5555)
	GenTree[string, *TestAAA, TreeEntity[string, *TestAAA]](nodes)
}
