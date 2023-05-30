package main

import "testing"

func TestNode_show(t *testing.T) {

	allAuthNode := make([]*FolderIsTree, 0)
	allAuthNode = append(allAuthNode, &FolderIsTree{name: "A"})
	allAuthNode = append(allAuthNode, &FolderIsTree{name: "B"})
	allAuthNode = append(allAuthNode, &FolderIsTree{name: "C"})
	allAuthNode = append(allAuthNode, &FolderIsTree{name: "D"})

	authNodes := make([]IData[string, *FolderIsTree], len(allAuthNode))

	for i, node := range allAuthNode {
		authNodes[i] = node
	}

	authTree := genTree[string, *FolderIsTree](authNodes)

	//authTree.DFS(func(node *Node[IData[string, *FolderIsTree]]) {
	authTree.DFS(func(node *Node[string, *FolderIsTree]) {
		node.data.showData()
	})

}
