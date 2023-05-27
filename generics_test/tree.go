package main

import "fmt"

type IComposite[T comparable] interface {
	GetID() string
	GetParentID() string
	SetParentID(string)
	Value() T
	AddChild(node IComposite[T])
	IsEqual(node IComposite[T]) bool
	Show()
}

type Node[T comparable] struct {
	id       string
	parentID string
	data     T
	children []IComposite[T]
}

func (n *Node[T]) GetID() string {
	return n.id
}

func (n *Node[T]) GetParentID() string {
	return n.parentID
}

func (n *Node[T]) SetParentID(id string) {
	n.parentID = id
}

func (n *Node[T]) Value() T {
	return n.data
}

func (n *Node[T]) AddChild(node IComposite[T]) {
	node.SetParentID(n.GetID())
	n.children = append(n.children, node)
}

func (n *Node[T]) IsEqual(node IComposite[T]) bool {
	r := false
	if n.GetID() == node.GetID() {
		r = true
	}
	return r
}

func (n *Node[T]) Show() {
	fmt.Printf("parentid: %v , id: %v\n", n.GetParentID(), n.GetID())
	for _, v := range n.children {
		v.Show()
	}
}

type Tree[T comparable] struct {
	root    IComposite[T]
	nodeMap map[string]IComposite[T]
}

type trees[T comparable] struct {
}

func (t trees[T]) AddChild(tree *Tree[T], parentNode IComposite[T], node IComposite[T]) {
	nodeID := node.GetID()

	if _, ok := tree.nodeMap[nodeID]; ok {
		panic("error root id " + nodeID)
	}
	if nil == tree.root {
		tree.root = parentNode
	}
	parentNode.AddChild(node)
	tree.nodeMap[nodeID] = node
}

func (t trees[T]) SetParentNode(tree *Tree[T], rootNode IComposite[T]) {
	tree.root = rootNode
}

func (t trees[T]) Show(tree *Tree[T]) {
	tree.root.Show()
}
