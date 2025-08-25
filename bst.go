package main

type BinarySearchTree struct {
	root *Node
}

type Node struct {
	value string
	left *Node
	right *Node
}

func (bst *BinarySearchTree) Insert(name string) {
	bst.root = insertNode(bst.root, name)
}

func insertNode(current *Node, name string) *Node {
	if current == nil {
		return &Node{value: name}
	}

	if name < current.value {
		current.left = insertNode(current.left, name)
	} else if name > current.value {
		current.right = insertNode(current.right, name)
	}

	return current
}
