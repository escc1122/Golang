package main

import (
	"strconv"
)

type MyGenericInterface[T Number] interface {
	DoSomething(T) string
}

type MyStruct struct {
}

func (s *MyStruct) DoSomething(i int) string {
	return "Received integer: " + strconv.Itoa(i)
}

// =================================================

// 研究中
type ComparableIface interface {
	IsEqual(a interface{}) bool
}

type TreeEntity[KeyType comparable, ValueType comparable] interface {
	GetID() KeyType
	GetParent() KeyType
	Value() ValueType
	ComparableIface
}

type Tree2[KeyType comparable, ValueType comparable] struct {
	root    *Node2[TreeEntity[KeyType, ValueType]]
	nodeMap map[KeyType]*Node2[TreeEntity[KeyType, ValueType]]
}

type Node2[T ComparableIface] struct {
	data     T
	children []*Node2[T]
}

func GenTree[KeyType comparable, ValueType comparable, T TreeEntity[KeyType, ValueType]](nodes []T) (*Tree2[KeyType, ValueType], error) {
	tree := &Tree2[KeyType, ValueType]{}
	tree.nodeMap = map[KeyType]*Node2[TreeEntity[KeyType, ValueType]]{}
	rootKey := *new(KeyType)
	tree.nodeMap[rootKey] = &Node2[TreeEntity[KeyType, ValueType]]{}
	tree.root = tree.nodeMap[rootKey]

	return tree, nil
}

type TestAAA struct {
	Id     string
	Name   string
	Parent string
	Type   int
	Sort   int
}

func (node *TestAAA) GetID() string {
	return node.Id
}

func (node *TestAAA) GetParent() string {
	return node.Parent
}

func (node *TestAAA) Value() *TestAAA {
	return node
}

func (node *TestAAA) IsEqual(a interface{}) bool {
	if str, ok := a.(string); ok {
		return str == node.Id
	}
	return false
}
