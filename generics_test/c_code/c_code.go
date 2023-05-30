package main

import "fmt"

type IData interface {
	showData()
}

type ITree[K comparable, V comparable] interface {
	GetID() K
	Value() V
	showData()
	//IData
}

type Tree[K comparable, V comparable] struct {
	root    *Node[ITree[K, V]]
	nodeMap map[K]*Node[ITree[K, V]]
}

type Node[T IData] struct {
	data     T
	children []*Node[T]
}

func genTree[K comparable, V comparable](nodes []ITree[K, V]) *Tree[K, V] {
	tree := &Tree[K, V]{}
	tree.nodeMap = map[K]*Node[ITree[K, V]]{}
	for _, node := range nodes {
		tree.nodeMap[node.GetID()] = &Node[ITree[K, V]]{data: node}
	}
	return tree
}

func (t *Tree[K, V]) DFS(callback func(*Node[ITree[K, V]])) {
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
