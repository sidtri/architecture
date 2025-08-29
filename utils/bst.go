package utils

type BinarySearchTree struct {
	Root *Node
}

func (bst *BinarySearchTree) Insert(name string) {
	bst.Root = insertNode(bst.Root, name)
}

func insertNode(current *Node, name string) *Node {
	if current == nil {
		return &Node{Value: name}
	}

	if name < current.Value {
		current.Left = insertNode(current.Left, name)
	} else if name > current.Value {
		current.Right = insertNode(current.Right, name)
	}

	return current
}
