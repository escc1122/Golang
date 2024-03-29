package main

import (
	"fmt"
)

// 重構用代碼
type ComparableIface interface {
	showData()
}

type IData[K comparable, V comparable] interface {
	GetID() K
	Value() V
	ComparableIface
}

type Tree[K comparable, V comparable] struct {
	root    *Node[K, V]
	nodeMap map[K]*Node[K, V]
}

type Node[K comparable, V comparable] struct {
	data     IData[K, V]
	children []*Node[K, V]
}

func genTree[K comparable, V comparable](nodes []IData[K, V]) *Tree[K, V] {
	tree := &Tree[K, V]{}
	tree.nodeMap = map[K]*Node[K, V]{}
	for _, node := range nodes {
		tree.nodeMap[node.GetID()] = &Node[K, V]{data: node}
	}
	return tree
}

func (t *Tree[K, V]) DFS(callback func(*Node[K, V])) {
	for _, v := range t.nodeMap {
		callback(v)
	}
}

type FolderIsTree struct {
	name string
}

func (f *FolderIsTree) showData() {
	fmt.Println(f.name)
}

func (f *FolderIsTree) GetID() string {
	return f.name
}

func (f *FolderIsTree) Value() *FolderIsTree {
	return f
}
