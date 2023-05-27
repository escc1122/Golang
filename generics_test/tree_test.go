package main

import (
	"testing"
)

func TestNode(t *testing.T) {
	type aaaaa struct {
		ID string
	}

	node2 := &Node[*aaaaa]{
		id: "2",
		data: &aaaaa{
			ID: "2",
		},
	}

	node3 := &Node[*aaaaa]{
		id: "3",
		data: &aaaaa{
			ID: "3",
		},
	}

	node4 := &Node[*aaaaa]{
		id: "4",
		data: &aaaaa{
			ID: "4",
		},
	}

	rootNode := &Node[*aaaaa]{
		id: "1",
		data: &aaaaa{
			ID: "1",
		},
	}

	rootNode.AddChild(node2)

	node2.AddChild(node3)

	node3.AddChild(node4)

	rootNode.Show()
}

func TestTree(t *testing.T) {
	type aaaaa struct {
		ID string
	}

	node2 := &Node[*aaaaa]{
		id: "2",
		data: &aaaaa{
			ID: "2",
		},
	}

	node3 := &Node[*aaaaa]{
		id: "3",
		data: &aaaaa{
			ID: "3",
		},
	}

	node4 := &Node[*aaaaa]{
		id: "4",
		data: &aaaaa{
			ID: "4",
		},
	}

	//var rootNode INode[*aaaaa]

	rootNode := &Node[*aaaaa]{
		id: "1",
		data: &aaaaa{
			ID: "1",
		},
	}

	tree := &Tree[*aaaaa]{
		nodeMap: make(map[string]IComposite[*aaaaa]),
	}

	aaa := trees[*aaaaa]{}

	aaa.AddChild(tree, rootNode, node2)

	aaa.AddChild(tree, node2, node3)

	aaa.AddChild(tree, node3, node4)

	aaa.Show(tree)
}
